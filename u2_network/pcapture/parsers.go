package main

import (
	"encoding/binary"
	"sort"
	"strings"
)

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
			PacketLength:          packet_length,
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
			EthType:   binary.BigEndian.Uint16(payload[12:14]), // all IPv4 (0x0800)
			Payload:   payload[14:],
		}
		all_frames = append(all_frames, eth_frame)
	}
	return all_frames
}

func parseIPDatagrams(eth_frames []EthernetFrame) []IPDatagram {
	all_datagrams := make([]IPDatagram, 0)
	for _, frame := range eth_frames {
		payload := frame.Payload
		datagram := IPDatagram{
			Version:        payload[0] >> 4,
			IHL:            payload[0] & 0x0f,
			DSCP:           payload[1] >> 2,
			ECN:            payload[1] & 0x03,
			TotalLength:    binary.BigEndian.Uint16(payload[2:4]),
			Identification: binary.BigEndian.Uint16(payload[4:6]),
			Flags:          payload[6] >> 5,
			FragmentOffset: binary.BigEndian.Uint16(payload[6:8]) & 0x1fff,
			TTL:            payload[8],
			Protocol:       payload[9], // all TCP baby
			HeaderChecksum: binary.BigEndian.Uint16(payload[10:12]),
			SourceIP:       binary.BigEndian.Uint32(payload[12:16]),
			DestinationIP:  binary.BigEndian.Uint32(payload[16:20]),
		}

		options_len := uint16(datagram.IHL-5) * 4

		if options_len > 0 {
			datagram.Options = payload[20 : 20+options_len]
		}
		payloadLength := datagram.TotalLength - uint16(datagram.IHL*4)
		datagram.Payload = payload[20+options_len : 20+options_len+payloadLength]
		all_datagrams = append(all_datagrams, datagram)
	}
	return all_datagrams
}

func parseTCPSegments(ip_datagrams []IPDatagram) []TCPSegment {
	all_segments := make([]TCPSegment, 0)
	for _, gram := range ip_datagrams {
		payload := gram.Payload
		segment := TCPSegment{
			SourcePort: binary.BigEndian.Uint16(payload[:2]),
			DestPort:   binary.BigEndian.Uint16(payload[2:4]),
			SequenceNo: binary.BigEndian.Uint32(payload[4:8]),
			AckNo:      binary.BigEndian.Uint32(payload[8:12]),
			DataOffset: (payload[12] >> 4) & 0x0f,
			CWR:        (payload[13]>>7)&0x01 == 1,
			ECE:        (payload[13]>>6)&0x01 == 1,
			URG:        (payload[13]>>5)&0x01 == 1,
			ACK:        (payload[13]>>4)&0x01 == 1,
			PSH:        (payload[13]>>3)&0x01 == 1,
			RST:        (payload[13]>>2)&0x01 == 1,
			SYN:        (payload[13]>>1)&0x01 == 1,
			FIN:        payload[13]&0x01 == 1,
			WindowSize: binary.BigEndian.Uint16(payload[14:16]),
			Checksum:   binary.BigEndian.Uint16(payload[16:18]),
			UrgPointer: binary.BigEndian.Uint16(payload[18:20]),
		}

		header_len := segment.DataOffset * 4
		options_len := header_len - 20
		if options_len > 0 {
			segment.Options = payload[20 : 20+options_len]
		}

		end_index := int(header_len)
		if end_index > len(payload) {
			end_index = len(payload)
		}

		segment.Payload = payload[end_index:]

		all_segments = append(all_segments, segment)

	}
	// sort by SequenceNo
	sort.Slice(all_segments, func(i, j int) bool {
		return all_segments[i].SequenceNo < all_segments[j].SequenceNo
	})

	return all_segments
}

func removeDuplicateSegments(tcp_segments []TCPSegment) []TCPSegment {
	uniqueSegments := make([]TCPSegment, 0, len(tcp_segments))
	seen := make(map[uint32]bool)

	for _, seg := range tcp_segments {
		if _, exists := seen[seg.SequenceNo]; !exists {
			seen[seg.SequenceNo] = true
			uniqueSegments = append(uniqueSegments, seg)
		}
	}
	return uniqueSegments
}

func repairHTTPData(orderedSegments []TCPSegment) string {
	allData := make([]string, len(orderedSegments))
	for _, seg := range orderedSegments {
		allData = append(allData, string(seg.Payload))
	}
	return strings.Join(allData, "")
}

func extractHTTPHeader(data string) (string, string) {
	split := strings.Split(data, "\r\n\r\n")
	return split[0], split[1]
}
