package skipList

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"underwater/WAL"
	"underwater/ssTable"
	"underwater/types"
	"underwater/utils"
)

type Database struct {
	Header *types.Node
	Level  int
	Size   int
	Tail   *types.Node
}

func (d *Database) Get(searchKey []byte) ([]byte, error) {
	current, _, match, err := d.search(searchKey)
	if err != nil {
		panic(err.Error())
	}
	if match {
		return current.Value, nil
	}
	msg := fmt.Sprintf("No entry for key %s found in database.", string(searchKey))
	return nil, errors.New(msg)
}

func (d *Database) Has(searchKey []byte) (bool, error) {
	_, _, match, err := d.search(searchKey)
	if err != nil {
		panic(err.Error())
	}
	if match {
		return true, nil
	}
	return false, nil
}

func (d *Database) Put(searchKey, newValue []byte) error {
	if d.Size >= utils.MEMTABLE_CAP {
		d.flushssTable()
	}
	current, update, match, err := d.search(searchKey)
	if err != nil {
		panic(err.Error())
	}

	wal := WAL.WAL{}
	entry, err := wal.CreateEntry(searchKey, newValue, utils.PUT)
	if err != nil {
		panic(err.Error())
	}
	wal.Write(entry)

	// check for and update an old key
	if match {
		d.Size -= len(current.Value)
		d.Size += len(newValue)
		current.Value = newValue
		return nil
	}

	// generate level, update DB level
	level := getLevel()
	if d.Level < level {
		for i := d.Level; i < level; i++ {
			update[i] = d.Header
		}
		d.Level = level
	}

	// create the actual node
	node := makeNode(searchKey, newValue)

	// fill the update list with our new node's pointer
	for i := 0; i < level; i++ {
		if update[i] == nil {
			continue
		}
		node.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = node

	}

	// Update table size for autoflushing purposes
	d.Size += len(searchKey) + len(newValue) + 4
	return nil
}

func (d *Database) Delete(searchKey []byte) error {
	current, update, match, err := d.search(searchKey)
	if err != nil {
		return err
	}
	wal := WAL.WAL{}
	entry, err := wal.CreateEntry(searchKey, nil, utils.DELETE)
	if err != nil {
		panic(err.Error())
	}
	err = wal.Write(entry)
	if err != nil {
		panic(err.Error())
	}

	// check for key
	if !match {
		// Didn't find the key to Delete
		error := fmt.Sprintf("Delete Error: No key %s in db", string(searchKey))
		return errors.New(error)
	}

	// repair the update list with the old node's forwards
	forward := current.Forward
	for i := 0; i < d.Level; i++ {
		if update[i].Forward[i] != current {
			continue
		}
		update[i].Forward[i] = forward[i]
	}

	// shrink list level
	for level := d.Level; level > 1; level-- {
		if d.Header.Forward[level-1] == nil {
			d.Level--
		}
	}
	return nil
}

func (d *Database) RangeScan(start, end []byte) (RSIterator, error) {
	head, _, match, err := d.search(start)
	if err != nil {
		return RSIterator{}, err
	}

	iter := RSIterator{}
	iter.entries = append(iter.entries, *head)
	current := head
	for {
		current = current.Forward[0]
		iter.entries = append(iter.entries, *current)
		if !utils.BLT(current.Forward[0].Key, end) {
			if match {
				iter.entries = append(iter.entries, *current.Forward[0])
			}
			return iter, err
		}
	}
}

func (d *Database) LoadCSV() error {
	println("LOADING CSV")
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

func (d *Database) search(searchKey []byte) (node *types.Node, forward types.Forward, match bool, err error) {
	update := types.Forward{}
	current := d.Header

	// start at the highest current level to save time
	for level := d.Level; level > 0; level-- {
		next := current.Forward[level-1]
		for utils.BLT(next.Key, searchKey) {
			current = next
			next = current.Forward[level-1]
		}
		update[level-1] = current
	}

	current = current.Forward[0]
	var directMatch bool
	if current != nil {
		directMatch = utils.BEQ(current.Key, searchKey)
	} else {
		directMatch = false
	}

	return current, update, directMatch, nil
}

func makeNode(searchKey, newValue []byte) *types.Node {
	node := &types.Node{
		Key:     searchKey,
		Value:   newValue,
		Forward: types.Forward{},
	}
	return node
}

func getLevel() int {
	i := 1
	for {
		if i == utils.MAX_LVL {
			return i
		}
		x := rand.Float32()
		if x < 0.5 {
			break
		}
		i++
	}
	return i
}

func CreateDatabase() (Database, error) {
	var entries []WAL.Entry
	walEntries, err := WAL.ReadWAL()
	if err != nil {
		panic(err.Error())
	}
	entries = append(entries, walEntries...)

	// Set up nil tail
	tail := &types.Node{
		Key:     []byte(utils.LAST_KEY),
		Value:   nil,
		Forward: types.Forward{},
	}

	// instantiate DB
	db := Database{
		Header: &types.Node{
			Key:     []byte("HEADER"),
			Forward: types.Forward{},
		},
		Level: 1,
		Size:  0,
	}

	// Point header to tail
	for i := 0; i < utils.MAX_LVL; i++ {
		db.Header.Forward[i] = tail
	}
	db.Tail = tail

	// Restore WAL
	if len(entries) > 0 {
		for _, entry := range entries {
			if entry.Op() == utils.PUT {
				// Put node
				db.Put(entry.Key(), entry.Value())
			}
			if entry.Op() == utils.DELETE {
				// Delete node
				db.Delete(entry.Key())
			}
		}
	}

	return db, nil
}

func (d *Database) Print() {
	for level := d.Level; level > 0; level-- {

		entry := d.Header.Forward[level-1]
		fmt.Printf("%d. Header", level)
		for entry != nil {
			fmt.Printf(" --> %s", string(entry.Key))
			entry = entry.Forward[level-1]
		}
		println()
	}
}

// return an ordered list of all entries
func (d *Database) ordered() []*types.Node {
	var entries []*types.Node
	first := d.Header.Forward[0]
	entries = append(entries, first)
	for current := first.Forward[0]; current != nil; current = current.Forward[0] {
		entries = append(entries, current)
	}
	return entries
}

func (d *Database) flushssTable() error {
	entries := d.ordered()
	err := ssTable.Flush(entries)
	if err != nil {
		return err
	}

	// reset memtable
	d.Size = 0
	d.Level = 1
	for i := 0; i < utils.MAX_LVL; i++ {
		d.Header.Forward[i] = d.Tail
	}

	// remove WAL
	os.Remove("WAL.log")
	println("ssTable file written!")
	return nil
}
