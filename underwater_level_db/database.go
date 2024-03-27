package main

import (
	"errors"
	"fmt"
	"sort"
)

type Database struct {
	entries []Entry
}

type Entry struct {
	key   []byte
	value []byte
}

func (d *Database) Get(key []byte) ([]byte, error) {
	for _, entry := range d.entries {
		if BEQ(key, entry.key) {
			return entry.value, nil
		}
	}
	error := fmt.Sprintf("Get Error: No key \"%s\" in db", string(key))
	return nil, errors.New(error)
}

func (d *Database) Has(key []byte) (bool, error) {
	for _, entry := range d.entries {
		if BEQ(key, entry.key) {
			return true, nil
		}
	}
	return false, nil
}

func (d *Database) Put(key, value []byte) error {
	for i, entry := range d.entries {
		if BEQ(key, entry.key) {
			d.entries[i].value = value
			return nil
		}
	}
	d.entries = append(d.entries, Entry{key, value})
	d.Sort()
	return nil
}

func (d *Database) Delete(key []byte) error {
	for i, entry := range d.entries {
		if BEQ(key, entry.key) {
			d.entries[i].value = nil
			return nil
		}
	}
	// Didn't find the key to Delete
	error := fmt.Sprintf("Delete Error: No key %s in db", string(key))
	return errors.New(error)
}

func (d *Database) RangeScan(start, limit []byte) (Iterator, error) {
	if BLT(limit, start) {
		limit, start = start, limit
	}
	iStart, iEnd := -1, -1
	iter := &RSIterator{}
	for i, entry := range d.entries {
		if (BLT(start, entry.key) || BEQ(start, entry.key)) && iStart < 0 {
			// We have our starting point
			iStart = i
		}
		if BLT(limit, entry.key) || i+1 == len(d.entries) {
			// We have our end and exit point
			iEnd = i - 1
			iter.entries = d.entries[iStart:iEnd]
			iter.index = 0
			return iter, nil
		}
	}

	// If we've gotten here, we never found a fitting entry and should return an Iterator with no entries
	return iter, nil
}

func (d *Database) Sort() {
	sort.Slice(d.entries, func(i, j int) bool {
		return string(d.entries[i].key) < string(d.entries[j].key)
	})
}
