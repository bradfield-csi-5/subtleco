package main

import (
	"golang.org/x/sys/unix"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// should be using listen, accept, recv, send

func main() {
	startServer()
}

func startServer() {
	// Create a socket
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	check(err)

	// Bind the socket
	ip := [4]byte{127, 0, 0, 1}
	port := 8080
	addr := &unix.SockaddrInet4{Port: port, Addr: ip}
	unix.Bind(fd, addr)

	// Listen on the socket
	unix.Listen(fd, 10)

	// Keep the socket listening
	for {
		// Accept a connection
		connFd, _, err := unix.Accept(fd)
		check(err)

		// Handle the connection
		handleConnection(connFd)

		// Close the connection
		unix.Close(connFd)
	}
}

func handleConnection(fd int) {
	nothingBuf := make([]byte, 512)
	message := "hello, child!"
	buf := []byte(message)

	for {
		unix.Recvfrom(fd, nothingBuf, 0)

		unix.Sendto(fd, buf[:13], 0, nil)
	}
}
