package main

import (
	"bytes"
	"encoding/binary"
	"strings"

	"golang.org/x/sys/unix"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type DNSHeader struct {
	ID      uint16
	Flags   uint16
	QDCount uint16
	ANCount uint16
	NSCound uint16
	ARCount uint16
}

type DNSQuestion struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

// helper to encode a domain name in DNS style
func encodeDomainName(domain string) []byte {
	var result []byte
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		result = append(result, byte(len(label)))
		result = append(result, label...)
	}
	result = append(result, 0)
	return result
}

func serializeHeader(header DNSHeader) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, header)
	return buf.Bytes()
}

func serializeQuestion(question DNSQuestion) []byte {
	buf := new(bytes.Buffer)
	buf.Write(question.QName)
	binary.Write(buf, binary.BigEndian, question.QType)
	binary.Write(buf, binary.BigEndian, question.QClass)
	return buf.Bytes()
}

// create a socket
func main() {
	socket, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	check(err)
	defer unix.Close(socket)

	// encode message
	header := DNSHeader{
		ID:      1234,
		Flags:   0x0100, // standard query, recursion desired
		QDCount: 1,
	}

	domain := "wikipedia.org"
	encodedDomain := encodeDomainName(domain)

	question := DNSQuestion{
		QName:  encodedDomain,
		QType:  1, // A record
		QClass: 1, // Internet
	}

	query := append(serializeHeader(header), serializeQuestion(question)...)

	// send message
	serverIP := [4]byte{8, 8, 8, 8}
	serverAddr := unix.SockaddrInet4{Addr: serverIP, Port: 53}

	err = unix.Sendto(socket, query, 0, &serverAddr)
	check(err)

	// recieve message
	buf := make([]byte, 512) // typical max size for a DNS message
	n, _, err := unix.Recvfrom(socket, buf, 0)
	check(err)

	print(string(buf[:n]))

	// decode message
}
