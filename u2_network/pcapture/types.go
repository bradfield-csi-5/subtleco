package main

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
	PacketLength          uint32
	FullPacketLength      uint32
	Payload               []byte
}

type EthernetFrame struct {
	MacDest   []byte
	MacSource []byte
	EthType   uint16
	Payload   []byte
}

type IPDatagram struct {
	Version        uint8
	IHL            uint8
	DSCP           uint8
	ECN            uint8
	TotalLength    uint16
	Identification uint16
	Flags          uint8
	FragmentOffset uint16
	TTL            uint8
	Protocol       uint8
	HeaderChecksum uint16
	SourceIP       uint32
	DestinationIP  uint32
	Options        []byte
	Payload        []byte
}

type TCPSegment struct {
	SourcePort uint16
	DestPort   uint16
	SequenceNo uint32
	AckNo      uint32
	DataOffset uint8
	CWR        bool
	ECE        bool
	URG        bool
	ACK        bool
	PSH        bool
	RST        bool
	SYN        bool
	FIN        bool
	WindowSize uint16
	Checksum   uint16
	UrgPointer uint16
	Options    []byte
	Payload    []byte
}
