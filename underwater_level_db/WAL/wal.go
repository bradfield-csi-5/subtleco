package WAL

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"underwater/utils"
)

type WAL struct {
	entries []Entry
}

type Entry interface {
	Op() byte
	Key() []byte
	KeyLen() uint16
	Value() []byte
	ValueLen() uint16
}

type DeleteEntry struct {
	op     byte
	keyLen uint16
	key    []byte
}

func (d *DeleteEntry) Op() byte {
	return d.op
}

func (d *DeleteEntry) Key() []byte {
	return d.key
}

func (d *DeleteEntry) KeyLen() uint16 {
	return d.keyLen
}

func (d *DeleteEntry) Value() []byte {
	return nil
}

func (d *DeleteEntry) ValueLen() uint16 {
	return 0
}

type PutEntry struct {
	DeleteEntry
	valueLen uint16
	value    []byte
}

func (p *PutEntry) Value() []byte {
	return p.value
}

func (p *PutEntry) ValueLen() uint16 {
	return p.valueLen
}

func (w *WAL) CreateEntry(key, value []byte, op byte) (Entry, error) {
	if op != utils.DELETE && op != utils.PUT {
		msg := fmt.Sprintf("Invalid op code for WAL: %b", op)
		return nil, errors.New(msg)
	}
	deleteEntry := DeleteEntry{
		op:     op,
		keyLen: uint16(len(key)),
		key:    key,
	}
	if op == utils.PUT {
		putEntry := PutEntry{
			DeleteEntry: deleteEntry,
			value:       value,
			valueLen:    uint16(len(value)),
		}
		return &putEntry, nil
	}
	return &deleteEntry, nil
}

func (w *WAL) Write(entry Entry) error {
	f, err := os.OpenFile(utils.LOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		panic(err.Error())
	}
	stream, err := w.byteStream(entry)
	if err != nil {
		panic(err.Error())
	}
	if _, err := f.Write(stream); err != nil {
		panic(err.Error())
	}
	return nil
}

func (w *WAL) byteStream(entry Entry) ([]byte, error) {
	var buf []byte
	buf = append(buf, entry.Op())

	// make sure keyLen is 2 bytes long regardless of contents
	keyLen := make([]byte, 2)
	binary.BigEndian.PutUint16(keyLen, entry.KeyLen())
	buf = append(buf, keyLen...)

	buf = append(buf, entry.Key()...)
	if putEntry, ok := entry.(*PutEntry); ok && entry.Op() == utils.PUT {

		// make sure valueLen is 2 bytes long regardless of contents
		valueLen := make([]byte, 2)
		binary.BigEndian.PutUint16(valueLen, putEntry.ValueLen())
		buf = append(buf, valueLen...)
		buf = append(buf, putEntry.Value()...)
	}
	return buf, nil
}

func ReadWAL() ([]Entry, error) {
	f, err := os.Open(utils.LOG)
	if err != nil {
		panic(err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err.Error())
	}

	var entries []Entry
	if fi.Size() > 0 {
		var deleteEntry DeleteEntry
		for {
			// Op Code
			opCode := make([]byte, 1)
			_, err := f.Read(opCode)
			if err == io.EOF {
				break
			}
			deleteEntry.op = opCode[0]

			// Key Len
			keyLen := make([]byte, 2)
			_, err = f.Read(keyLen)
			if err != nil {
				panic(err.Error())
			}
			deleteEntry.keyLen = uint16(binary.BigEndian.Uint16(keyLen))

			// Key
			key := make([]byte, deleteEntry.keyLen)
			_, err = f.Read(key)
			deleteEntry.key = key

			if opCode[0] == utils.PUT {
				putEntry := PutEntry{DeleteEntry: deleteEntry}
				// Value Len
				valueLen := make([]byte, 2)
				_, err = f.Read(valueLen)
				if err != nil {
					panic(err.Error())
				}
				putEntry.valueLen = uint16(binary.BigEndian.Uint16(valueLen))

				// Key
				value := make([]byte, putEntry.valueLen)
				_, err = f.Read(value)
				if err != nil {
					panic(err.Error())
				}
				putEntry.value = value

				entries = append(entries, &putEntry)
			} else {
				entries = append(entries, &deleteEntry)
			}
		}
	}
	return entries, nil
}
