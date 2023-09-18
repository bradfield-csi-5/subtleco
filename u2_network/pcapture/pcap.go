package main

import (
	"os"
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

	// Parse Pcap Packets
	filePayloadSize := fileSize - 24
	pcapPackets := splitPcapPackets(pcacFile.Payload, filePayloadSize)

	// Parse Ethernet Frames
	ethFrames := parseEthernetFrames(pcapPackets)

	// Parse IP Datagrams
	ip_datagrams := parseIPDatagrams(ethFrames)

	// Parse TCP Segments
	tcpSegments := parseTCPSegments(ip_datagrams)
	uniqueSegments := removeDuplicateSegments(tcpSegments)

	// Stitch back together
	httpData := repairHTTPData(uniqueSegments)

	// Extract body
	_, body := extractHTTPHeader(httpData)

	// Write to file
	os.WriteFile("dude.jpeg", []byte(body), 0666)
}
