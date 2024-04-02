package skipList

import (
	"math/rand"
	"underwater/utils"
)

const (
	MAX_LVL  = 10
	LAST_KEY = "zzzzzzzzzz"
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

func (d *Database) Put(searchKey, newValue []byte) error {
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

	// check for and update an old key
	if utils.BEQ(current.Key, searchKey) {
		current.Value = newValue
	} else {
		// create a new node, updating the DB.Level if needed
		level := getLevel()
		if d.Level < level {
			for i := d.Level; i <= level; i++ {
				update[i] = d.Header
			}
			d.Level = level
		}
		node := makeNode(searchKey, newValue)

		for i := 0; i < level; i++ {
			if update[i] == nil {
				continue
			}
			node.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = node
		}
	}
	return nil
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

func (d *Database) Get(key []byte) ([]byte, error) {
	return []byte(""), nil
}

func (d *Database) Has(key []byte) (bool, error) {
	return false, nil
}

func (d *Database) Delete(key []byte) error {
	return nil
}
