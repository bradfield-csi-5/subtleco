package main

import (
	"fmt"
	"underwater/skipList"
)

func main() {
	db, err := skipList.CreateDatabase()
	if err != nil {
		panic(err.Error())
	}

	db.Put([]byte("pizza"), []byte("yummo"))
	db.Put([]byte("goblin"), []byte("ickko"))
	db.Put([]byte("water"), []byte("fine"))
	db.Put([]byte("goblin"), []byte("ickkYYYYYYYYY"))
	db.Put([]byte("goldfish"), []byte("xyz"))
	db.Put([]byte("darren"), []byte("xyz"))
	db.Put([]byte("apple"), []byte("xyz"))

	for level := db.Level; level > 0; level-- {

		entry := db.Header.Forward[level-1]
		fmt.Printf("%d. Header", level)
		for entry != nil {
			fmt.Printf(" --> %s", string(entry.Key))
			entry = entry.Forward[level-1]
		}
		println()
	}
}
