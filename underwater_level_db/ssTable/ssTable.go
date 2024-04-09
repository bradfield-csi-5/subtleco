package ssTable

import (
	"encoding/binary"
	"os"
	"time"
	"underwater/types"
)

type SSTable struct {
	directoryOffset uint16
	entries         []Entry
	directory       []DirEntry
}

func (t *SSTable) streamData() []byte {
	// make the byte array
	var data []byte

	// start with the directory offset
	data = binary.BigEndian.AppendUint16(data, t.directoryOffset)

	// then stream in all of the entries
	for _, entry := range t.entries {
		data = binary.BigEndian.AppendUint16(data, entry.keyLen)
		data = append(data, entry.key...)
		data = binary.BigEndian.AppendUint16(data, entry.valueLen)
		data = append(data, entry.value...)
	}

	// finally, append the directory. Note - despite previous plans this will just go in order at the end of the entries
	for _, dirEntry := range t.directory {
		data = binary.BigEndian.AppendUint16(data, dirEntry.KeyLen)
		data = append(data, dirEntry.Key...)
		data = binary.BigEndian.AppendUint16(data, dirEntry.Offset)
	}
	return data
}

type Entry struct {
	keyLen   uint16
	key      []byte
	valueLen uint16
	value    []byte
}

type DirEntry struct {
	Offset uint16
	Key    []byte
	KeyLen uint16
}

func Flush(entries []*types.Node) error {
	var directory []DirEntry
	top := 2 // leave room for directory offset of 2 bytes
	for i, entry := range entries {
		if entry.Forward[0] == nil {
			break
		}
		// Only add every 10 things to the directory
		if i%10 == 0 {
			dirEntry := DirEntry{
				Offset: uint16(top),
				Key:    entry.Key,
				KeyLen: uint16(len(entry.Key)),
			}
			directory = append(directory, dirEntry)
		}
		entryLen := len(entry.Key) + len(entry.Value) + 4 // 4 is padding for two uint16s - key length and value length
		top += entryLen
	}
	directoryOffset := top

	err := writeSSTable(directoryOffset, entries, directory)
	if err != nil {
		return err
	}
	return nil
}

func writeSSTable(dirOff int, nodes []*types.Node, directory []DirEntry) error {
	var entries []Entry
	for _, node := range nodes {
		if node.Forward[0] != nil {
			entry := Entry{
				keyLen:   uint16(len(node.Key)),
				key:      node.Key,
				valueLen: uint16(len(node.Value)),
				value:    node.Value,
			}
			entries = append(entries, entry)
		}
	}
	table := SSTable{
		directoryOffset: uint16(dirOff),
		entries:         entries,
		directory:       directory,
	}
	tableStream := table.streamData()
	fileName := getFilenameFromDatetime()
	err := os.WriteFile(fileName, tableStream, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getFilenameFromDatetime() string {
	now := time.Now()
	fileName := now.Format("20060102-150405") + ".sst"
	return fileName
}
