package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	PCAP_FILE_HEADER   = 24
	PCAP_PACKET_HEADER = 16
)

type PcapFile struct {
	MagicNumber    uint32
	VersionMajor   uint16
	VersionMinor   uint16
	TimeZoneOffset uint32
	Accuracy       uint32
	SnapshotLength uint32
	LinkHeaderType uint32
	Payload        []byte
}

type PcapPacket struct {
	TimeStampSeconds      uint32
	TimeStampMicroseconds uint32
	PacketLenght          uint32
	FullPacketLength      uint32
	Payload               []byte
}

type EthernetFrame struct {
	MacDest   []byte
	MacSource []byte
	EthType   uint16
	Payload   []byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// crack open that file
	file, err := os.ReadFile("net.cap")
	check(err)
	file_size := len(file)

	// copy the contents into a properly sized buffer
	buf := make([]byte, file_size)
	copy(buf, file)
	check(err)

	// Get file header
	pcap_file := parsePcapFile(buf)

	// Split into Pcap Packets
	file_payload_size := file_size - PCAP_FILE_HEADER
	pcap_packets := splitPcapPackets(pcap_file.Payload, file_payload_size)
	fmt.Println(len(pcap_packets))

	// Split into Ethernet Frames
	eth_frames := parseEthernetFrames(pcap_packets)
	fmt.Println(len(eth_frames))

	// Parse IP Datagrams

	// Parse TCP Segments

	// Stitch back together

	// Parse HTTP elements
}

func parsePcapFile(file []byte) PcapFile {
	return PcapFile{
		MagicNumber:    binary.LittleEndian.Uint32(file[:4]),
		VersionMajor:   binary.LittleEndian.Uint16(file[4:6]),
		VersionMinor:   binary.LittleEndian.Uint16(file[6:8]),
		TimeZoneOffset: binary.LittleEndian.Uint32(file[8:12]),
		Accuracy:       binary.LittleEndian.Uint32(file[12:16]),
		SnapshotLength: binary.LittleEndian.Uint32(file[16:20]),
		LinkHeaderType: binary.LittleEndian.Uint32(file[20:24]),
		Payload:        file[24:],
	}
}

func splitPcapPackets(packets []byte, file_size int) []PcapPacket {
	all_packets := make([]PcapPacket, 0)

	for packet_start := 0; packet_start < file_size; {
		packet_length := binary.LittleEndian.Uint32(packets[packet_start+8 : packet_start+12])
		pcap_packet := PcapPacket{
			TimeStampSeconds:      binary.LittleEndian.Uint32(packets[packet_start : packet_start+4]),
			TimeStampMicroseconds: binary.LittleEndian.Uint32(packets[packet_start+4 : packet_start+8]),
			PacketLenght:          packet_length,
			FullPacketLength:      binary.LittleEndian.Uint32(packets[packet_start+12 : packet_start+16]),
			Payload:               packets[packet_start+16 : packet_start+16+int(packet_length)],
		}

		all_packets = append(all_packets, pcap_packet)
		packet_start += PCAP_PACKET_HEADER + int(packet_length)
	}

	return all_packets
}

func parseEthernetFrames(packets []PcapPacket) []EthernetFrame {
	all_frames := make([]EthernetFrame, 0)
	for _, packet := range packets {
		payload := packet.Payload
		eth_frame := EthernetFrame{
			MacDest:   payload[:6],
			MacSource: payload[6:12],
			EthType:   binary.LittleEndian.Uint16(payload[12:16]),
			Payload:   payload[14:],
		}
		all_frames = append(all_frames, eth_frame)
	}
	return all_frames
}
