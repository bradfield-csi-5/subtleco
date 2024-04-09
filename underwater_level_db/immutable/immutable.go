package immutable

import (
	"encoding/binary"
	"fmt"
	"os"
)

type ImmuDB struct {
	filename string
	dir      Directory
}

type Directory struct {
	sparseKeys [][]byte
	offsets    []uint16
}

func (d *ImmuDB) Get(key []byte) ([]byte, error) {
	return []byte(""), nil
}

func (d *ImmuDB) Has(key []byte) (bool, error) {
	return false, nil
}
func (d *ImmuDB) RangeScan() {}

func CreateImmuDB(filename string) (ImmuDB, error) {
	immuDB := ImmuDB{}
	directoryStream, err := extractDirectoryStream(filename)
	if err != nil {
		return immuDB, err
	}
	directory, err := parseDirectoryStream(directoryStream)
	if err != nil {
		return immuDB, err
	}
	immuDB.filename = filename
	immuDB.dir = directory

	fmt.Printf("Filename: %s\n", filename)
	println("-------Directory-------")
	for i, entry := range immuDB.dir.sparseKeys {
		fmt.Printf("%s is at offset %x\n", string(entry), immuDB.dir.offsets[i])
	}
	return immuDB, err
}

func extractDirectoryStream(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, _ := f.Stat()
	fileLen := fi.Size()

	dirOffsetBuf := make([]byte, 2)
	f.Read(dirOffsetBuf)
	dirOffset := int64(binary.BigEndian.Uint16(dirOffsetBuf))

	dirLen := fileLen - dirOffset

	dirBuf := make([]byte, dirLen)

	f.Seek(dirOffset, 0)
	f.Read(dirBuf)
	return dirBuf, nil
}

func parseDirectoryStream(stream []byte) (Directory, error) {
	directory := Directory{}
	for i := 0; i < len(stream); {
		keyLenBuf := stream[i : i+2]
		i += 2
		keyLen := int(binary.BigEndian.Uint16(keyLenBuf))
		key := stream[i : i+keyLen]
		i += int(keyLen)
		offsetBuf := stream[i : i+2]
		offset := binary.BigEndian.Uint16(offsetBuf)
		i += 2

		directory.sparseKeys = append(directory.sparseKeys, key)
		directory.offsets = append(directory.offsets, offset)
	}

	return directory, nil
}
