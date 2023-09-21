package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

const HEADER_LEN = 12

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
	NSCount uint16
	ARCount uint16
}

type DNSQuestion struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

type ResourceRecord struct {
	Name  []byte
	Type  uint16
	Class uint16
	TTL   uint32
	RDLen uint16
	RData []byte
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

func decodeHeader(data []byte) DNSHeader {
	var header DNSHeader
	buf := bytes.NewReader(data)
	binary.Read(buf, binary.BigEndian, &header)
	return header
}

func decodeDomainName(data []byte, offset int) (string, int) {
	var nameParts []string
	OGoffset := 0
	pointer := false

	for {
		// Check for pointer
		if (data[offset] & 0xC0) == 0xC0 {
			if !pointer {
				OGoffset = offset + 2 // 21
			}
			fmt.Println("OG Offset: ", OGoffset)
			pointer = true
			ptr_offset := (data[offset]&0x3F)<<8 | data[offset+1] // 12
			offset = int(ptr_offset)
			continue
		}

		length := int(data[offset])
		if length == 0 {
			break
		}
		nameParts = append(nameParts, string(data[offset+1:offset+1+length]))
		offset += length + 1
	}

	if pointer {
		return strings.Join(nameParts, "."), OGoffset
	}

	return strings.Join(nameParts, "."), offset + 1
}

func decodeQuestion(data []byte) (DNSQuestion, int) {
	data = data[12:]
	name, offset := decodeDomainName(data, 0)
	qType := binary.BigEndian.Uint16(data[offset : offset+2])
	qClass := binary.BigEndian.Uint16(data[offset+2 : offset+4])
	return DNSQuestion{
		QName:  []byte(name),
		QType:  qType,
		QClass: qClass,
	}, offset + 4 + 12
}

func decodeResourceRecord(data []byte, offset int) (ResourceRecord, int) {
	name, offset := decodeDomainName(data, offset)
	rType := binary.BigEndian.Uint16(data[offset : offset+2])
	rClass := binary.BigEndian.Uint16(data[offset+2 : offset+4])
	ttl := binary.BigEndian.Uint32(data[offset+4 : offset+8])
	rDlen := binary.BigEndian.Uint16(data[offset+8 : offset+10])
	rData := data[offset+10 : offset+10+int(rDlen)]
	return ResourceRecord{
		Name:  []byte(name),
		Type:  rType,
		Class: rClass,
		TTL:   ttl,
		RDLen: rDlen,
		RData: rData,
	}, offset + 10 + int(rDlen)
}

func decodeResponse(data []byte) (DNSHeader, DNSQuestion, ResourceRecord) {
	header := decodeHeader(data[:12])
	question, offset := decodeQuestion(data)
	answer, _ := decodeResourceRecord(data, offset)

	return header, question, answer
}

func main() {
	// create a socket

	socket, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	check(err)
	defer unix.Close(socket)

	// encode message

	header := DNSHeader{
		ID:      1234,
		Flags:   0x0100, // standard query, recursion desired
		QDCount: 1,
	}

	domain := os.Args[1]
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

	// decode message
	_, _, answer := decodeResponse(buf[:n])
	destinationIP := fmt.Sprintf("%d.%d.%d.%d\n", answer.RData[0], answer.RData[1], answer.RData[2], answer.RData[3])

	fmt.Printf("The IP address for %s is %s", os.Args[1], destinationIP)
}
