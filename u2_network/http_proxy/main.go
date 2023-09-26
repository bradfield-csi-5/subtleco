package main

import (
	"fmt"
	"strings"

	"golang.org/x/sys/unix"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(" - Starting the server")
	// Create a cache
	cache := make(map[string][]byte)

	// Create a socket & bind it

	fmt.Println(" - Creating a listening socket")
	fd := createListener()

	// Keep the party going
	fmt.Println(" - Entering listening loop")
	for {
		// Accept a connection
		clientFd := acceptConnection(fd)

		// Receive a packet
		clientBuf := make([]byte, 512)
		_ = receivePacket(clientFd, &clientBuf)

		// Resolve cache
		cacheHit := resolveFromCache(clientFd, clientBuf, cache)

		if !cacheHit {
			resolveFromServer(clientFd, clientBuf, cache)
		}

		// Close the connection
		unix.Close(clientFd)
	}
}

func createListener() int {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	check(err)

	ip := [4]byte{127, 0, 0, 1}
	port := 8081
	addr := &unix.SockaddrInet4{Port: port, Addr: ip}
	err = unix.Bind(fd, addr)
	check(err)

	// Listen
	err = unix.Listen(fd, 10)
	check(err)
	return fd
}

func acceptConnection(fd int) int {
	connFd, _, err := unix.Accept(fd)
	check(err)
	return connFd
}

func receivePacket(fd int, buf *[]byte) int {
	n, _, err := unix.Recvfrom(fd, *buf, 0)
	if err != nil {
		fmt.Println("Error during Recvfrom:", err)
	}
	if n == 0 {
		fmt.Println("I didn't see anything from the client")
		fmt.Println("Connection closed by server.")
	}
	return n
}

func resolveFromServer(clientFd int, clientBuf []byte, cache map[string][]byte) {
	fmt.Println("**************")
	fmt.Println("* CACHE MISS *")
	fmt.Println("**************")
	fmt.Println()

	// Create new socket to reach out to server
	serverFd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	check(err)

	ip := [4]byte{127, 0, 0, 1}

	serverPort := 9000
	serverAddr := &unix.SockaddrInet4{Port: serverPort, Addr: ip}
	unix.Connect(serverFd, serverAddr)

	buf := make([]byte, 512)

	err = unix.Sendto(serverFd, clientBuf, 0, nil)
	if err != nil {
		fmt.Println("Error during Sendto:", err)
		return
	}

	serverBuf := make([]byte, 0)

	for {
		// Receive server response in chunks and relay it to the client
		n, _, err := unix.Recvfrom(serverFd, buf, 0)
		if err != nil {
			fmt.Println("Error during server side Recvfrom:", err)
			break
		}

		if n == 0 {

			// Add to cache
			parsedRequest := string(clientBuf)
			path := strings.Split(parsedRequest, " ")[1]

			fmt.Printf("Adding '%s' to cache....\n", path)
			cache[path] = serverBuf

			break
		}

		serverBuf = append(serverBuf, buf...)

		err = unix.Sendto(clientFd, buf[:n], 0, nil)
		if err != nil {
			fmt.Println("Error during Sendto:", err)
			break
		}

	}
	unix.Close(serverFd)
}

func resolveFromCache(clientFd int, clientBuf []byte, cache map[string][]byte) bool {
	parsedRequest := string(clientBuf)
	path := strings.Split(parsedRequest, " ")[1]
	res, found := cache[path]

	if found {

		fmt.Println("**************")
		fmt.Println("* CACHE HIT! *")
		fmt.Println("**************")
		fmt.Println()

		err := unix.Sendto(clientFd, res, 0, nil)
		if err != nil {
			fmt.Println("Error during Sendto:", err)
			return false
		}
	}
	return found
}
