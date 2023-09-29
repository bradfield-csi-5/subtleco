package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/unix"
)

type cyberData struct {
	Length   uint16
	Checksum uint16
	Corrupt  uint8
	Body     []byte
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// CLIENT
func main() {
	// Addresses
	ip := [4]byte{127, 0, 0, 1}
	clientPort := 9998
	proxyPort := 53874
	clientAddr := &unix.SockaddrInet4{Port: clientPort, Addr: ip}
	proxyAddr := &unix.SockaddrInet4{Port: proxyPort, Addr: ip}

	// Create a socket
	socket, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	check(err)
	defer unix.Close(socket)

	// Bind the socket
	unix.Bind(socket, clientAddr)

	for {

		// Encode a message
		encodedMessage := encodeMessage([]byte("✅✅✅\n"), 0)

		// Send message to server
		fmt.Println("*** Sending a message.........")
		fmt.Println()
		err = unix.Sendto(socket, encodedMessage, 0, proxyAddr)
		check(err)

		// Receive a reply
		buf := make([]byte, 512)
		_, _, err := unix.Recvfrom(socket, buf, 0)
		check(err)

		serverReply := decodeMessage(buf)
		if serverReply.Corrupt == 0 && serverReply.Checksum == generateChecksum(serverReply.Body) {
			fmt.Println("GREAT JOB EVERYONE")
			fmt.Println("*** Server responded with: ")
			fmt.Println(string(serverReply.Body))
			break
		}

		fmt.Println("❌❌❌")
		fmt.Println("Shit it got bad and we'll try again.....")
		fmt.Println()

	}
}

func decodeMessage(incoming []byte) cyberData {
	var message cyberData
	buf := bytes.NewReader(incoming)
	err := binary.Read(buf, binary.BigEndian, &message.Length)
	check(err)
	err = binary.Read(buf, binary.BigEndian, &message.Checksum)
	check(err)
	err = binary.Read(buf, binary.BigEndian, &message.Corrupt)
	check(err)

	message.Body = make([]byte, message.Length)
	_, err = buf.Read(message.Body)
	check(err)

	return message
}

func generateChecksum(message []byte) uint16 {
	acc := uint32(0)
	for _, b := range message {
		acc += uint32(b)
		carry := acc >> 16
		acc += carry
	}
	return uint16(acc)
}

func encodeMessage(body []byte, corrupt uint8) []byte {
	length := uint16(len(body))
	checksum := generateChecksum(body)

	message := cyberData{
		Length:   length,
		Checksum: checksum,
		Corrupt:  corrupt,
		Body:     body,
	}

	messageBuf := new(bytes.Buffer)

	// Write Length and Checksum first
	binary.Write(messageBuf, binary.BigEndian, message.Length)
	binary.Write(messageBuf, binary.BigEndian, message.Checksum)
	binary.Write(messageBuf, binary.BigEndian, message.Corrupt)

	// Append Body to the buffer
	messageBuf.Write(message.Body)

	return messageBuf.Bytes()
}
