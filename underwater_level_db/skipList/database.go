package skipList

import (
	"errors"
	"fmt"
	"math/rand"
	"underwater/utils"
)

const (
	MAX_LVL  = 10
	LAST_KEY = "zzzzzz"
)

type Forward [MAX_LVL]*Node

type Database struct {
	Header *Node
	Level  int
}

type Node struct {
	Key     []byte
	Value   []byte
	Forward Forward
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
	current, update, match, err := d.search(searchKey)
	if err != nil {
		panic(err.Error())
	}

	// check for and update an old key
	if match {
		current.Value = newValue
		return nil
	}

	// generate level, update DB level
	level := getLevel()
	if d.Level < level {
		for i := d.Level; i <= level; i++ {
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
	return nil
}

func (d *Database) Delete(searchKey []byte) error {
	current, update, match, err := d.search(searchKey)
	if err != nil {
		return err
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
	head, _, _, err := d.search(start)
	if err != nil {
		return RSIterator{}, err
	}
	tail, _, _, err := d.search(end)
	if err != nil {
		return RSIterator{}, err
	}

	iter := RSIterator{}
	iter.entries = append(iter.entries, *head)
	current := head
	for {
		current = current.Forward[0]
		iter.entries = append(iter.entries, *current)
		if current.Forward[0] == tail {
			if utils.BEQ(end, tail.Key) {
				iter.entries = append(iter.entries, *tail)
			}
			return iter, err
		}
	}
}

func (d *Database) search(searchKey []byte) (node *Node, forward Forward, match bool, err error) {
	update := Forward{}
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
	directMatch := utils.BEQ(current.Key, searchKey)

	return current, update, directMatch, nil
}

func makeNode(searchKey, newValue []byte) *Node {
	node := &Node{
		Key:     searchKey,
		Value:   newValue,
		Forward: Forward{},
	}
	return node
}

func getLevel() int {
	i := 1
	for {
		if i == MAX_LVL {
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
	// Set up nil tail
	tail := &Node{
		Key:     []byte(LAST_KEY),
		Value:   nil,
		Forward: Forward{},
	}

	// instantiate DB
	db := Database{
		Header: &Node{
			Key:     []byte("HEADER"),
			Forward: Forward{},
		},
		Level: 1,
	}

	// Point header to tail
	for i := 0; i < MAX_LVL; i++ {
		db.Header.Forward[i] = tail
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
