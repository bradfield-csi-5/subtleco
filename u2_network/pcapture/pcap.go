package main

import (
	"os"
)

const (
	PCAP_FILE_HEADER   = 24
	PCAP_PACKET_HEADER = 16
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// crack open that file
	file, err := os.ReadFile("net.cap")
	check(err)
	fileSize := len(file)

	// copy the contents into a properly sized buffer
	buf := make([]byte, fileSize)
	copy(buf, file)
	check(err)

	// Get file header
	pcacFile := parsePcapFile(buf)

	// Split into Pcap Packets
	filePayloadSize := fileSize - PCAP_FILE_HEADER
	pcapPackets := splitPcapPackets(pcacFile.Payload, filePayloadSize)

	// Split into Ethernet Frames
	ethFrames := parseEthernetFrames(pcapPackets)

	// Parse IP Datagrams
	ip_datagrams := parseIPDatagrams(ethFrames)

	// Parse TCP Segments
	tcpSegments := parseTCPSegments(ip_datagrams)
	uniqueSegments := removeDuplicateSegments(tcpSegments)

	// Stitch back together
	httpData := repairHTTPData(uniqueSegments)
	_, body := extractHTTPHeader(httpData)

	// Parse HTTP elements
	os.WriteFile("dude.jpeg", []byte(body), 0666)
}
