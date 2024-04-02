package simpleList

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"underwater/interfaces"
	"underwater/utils"
)

type Database struct {
	Entries []Entry
}

type Entry struct {
	key   []byte
	value []byte
}

func (d *Database) Get(key []byte) ([]byte, error) {
	for _, entry := range d.Entries {
		if utils.BEQ(key, entry.key) {
			return entry.value, nil
		}
	}
	error := fmt.Sprintf("Get Error: No key \"%s\" in db", string(key))
	return nil, errors.New(error)
}

func (d *Database) Has(key []byte) (bool, error) {
	for _, entry := range d.Entries {
		if utils.BEQ(key, entry.key) {
			return true, nil
		}
	}
	return false, nil
}

func (d *Database) Put(key, value []byte) error {
	for i, entry := range d.Entries {
		if utils.BEQ(key, entry.key) {
			d.Entries[i].value = value
			return nil
		}
	}
	d.Entries = append(d.Entries, Entry{key, value})
	d.Sort()
	return nil
}

func (d *Database) Delete(key []byte) error {
	for i, entry := range d.Entries {
		if utils.BEQ(key, entry.key) {
			d.Entries[i].value = nil
			return nil
		}
	}
	// Didn't find the key to Delete
	error := fmt.Sprintf("Delete Error: No key %s in db", string(key))
	return errors.New(error)
}

func (d *Database) RangeScan(start, limit []byte) (interfaces.Iterator, error) {
	if utils.BLT(limit, start) {
		limit, start = start, limit
	}
	iStart, iEnd := -1, -1
	iter := &RSIterator{}
	for i, entry := range d.Entries {
		if (utils.BLT(start, entry.key) || utils.BEQ(start, entry.key)) && iStart < 0 {
			// We have our starting point
			iStart = i
		}
		if utils.BLT(limit, entry.key) || i+1 == len(d.Entries) {
			// We have our end and exit point
			iEnd = i - 1
			iter.entries = d.Entries[iStart:iEnd]
			iter.index = 0
			return iter, nil
		}
	}

	// If we've gotten here, we never found a fitting entry and should return an Iterator with no entries
	return iter, nil
}

func (d *Database) Sort() {
	sort.Slice(d.Entries, func(i, j int) bool {
		return string(d.Entries[i].key) < string(d.Entries[j].key)
	})
}

func (d *Database) LoadCSV() error {
	file, err := os.Open("movies.csv")
	if err != nil {
		return errors.New("Failed to open file")
	}

	defer file.Close()

	// skip the header row
	csvReader := csv.NewReader(file)
	_, err = csvReader.Read()
	if err != nil {
		return errors.New("Failed to parse CSV")
	}

	for {
		item, err := csvReader.Read()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}

		d.Put([]byte(item[1]), []byte(item[2]))
	}
	return nil
}
