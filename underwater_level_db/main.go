package main

import (
	"fmt"
	"underwater/skipList"
	"underwater/utils"
)

func main() {
	db, err := skipList.CreateDatabase()
	if err != nil {
		panic(err.Error())
	}
	db.Print()
}

func main2() {
	db, err := skipList.CreateDatabase()
	if err != nil {
		panic(err.Error())
	}

	db.Put([]byte("pizzas"), []byte("yummo"))
	db.Put([]byte("goblin"), []byte("ickko"))
	db.Put([]byte("watery"), []byte("fine"))
	db.Put([]byte("gldfsh"), []byte("xyz"))
	db.Put([]byte("darren"), []byte("xyz"))
	db.Put([]byte("appler"), []byte("xyz"))

	// test edit
	preEdit, _ := db.Get([]byte("goblin"))
	db.Put([]byte("goblin"), []byte("ickkYYYYYYYYY"))
	postEdit, _ := db.Get([]byte("goblin"))
	if !utils.BEQ(preEdit, postEdit) {
		println("editing a key works")
	}

	// get success
	get, err := db.Get([]byte("watery"))
	if err == nil {
		println("Value for watery:", string(get))
	}

	// get fail
	_, err = db.Get([]byte("nothing"))
	if err != nil {
		println(err.Error())
	}

	// has
	has, _ := db.Has([]byte("watery"))
	if has {
		println("DB does have entry with key 'watery'")
	}

	// has not
	hasNot, _ := db.Has([]byte("buffalo"))
	if !hasNot {
		println("DB does NOT have entry with key 'buffalo'")
	}

	// Delete test
	key := "goblin"
	db.Delete([]byte(key))
	has, _ = db.Has([]byte(key))
	if !has {
		println("Successfully deleted", key)
	}

	// RangeScan test
	start := []byte("appler")
	end := []byte("pizzas")
	iter, err := db.RangeScan(start, end)
	fmt.Printf("\nRangeScan of %s to %s\n", string(start), string(end))
	println("-------------------------")
	fmt.Printf("%s ", string(iter.Key()))
	for iter.Next() {
		fmt.Printf("--> %s ", string(iter.Key()))
	}
	println()

	db.Print()
}
