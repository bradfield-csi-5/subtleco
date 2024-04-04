package main

import (
	"fmt"
	"strconv"
	"testing"
	"underwater/simpleList"
	"underwater/skipList"
)

const (
	size = 10000
	low  = "key0001"
	med  = "key5000"
	high = "key9999"
)

var (
	skipDB   skipList.Database
	simpleDB simpleList.Database
)

func init() {
	println("Initializing test DBs")
	skipDB, _ = skipList.CreateDatabase()
	simpleDB = simpleList.Database{}

	// Determine the width for zero padding based on the size variable.
	// This will ensure all keys have the same width.
	width := len(strconv.Itoa(size - 1))

	for i := 0; i < size; i++ {
		// Use fmt.Sprintf to format the key with leading zeros based on the width calculated.
		key := []byte(fmt.Sprintf("key%0*d", width, i))
		value := []byte(fmt.Sprintf("value%0*d", width, i))
		simpleDB.Put(key, value)
		skipDB.Put(key, value)
	}
	println("Initialization complete")
}

func BenchmarkSkipGet3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := []string{low, med, high}
		for _, key := range keys {
			testKey := []byte(key)
			_, err := skipDB.Get(testKey)
			if err != nil {
				b.Fatalf("Failed to get key %s: %v", testKey, err)
			}
		}
	}
}

func BenchmarkSimpleHas3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := []string{low, med, high}
		for _, key := range keys {
			testKey := []byte(key)
			has, err := simpleDB.Has(testKey)
			if err != nil {
				b.Fatalf("Has failed on key %s: %v", testKey, err)
			}
			if !has {
				b.Fatalf("Has reported key %s not present", testKey)
			}
		}
	}
}

func BenchmarkSkipHas3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := []string{low, med, high}
		for _, key := range keys {
			testKey := []byte(key)
			has, err := skipDB.Has(testKey)
			if err != nil {
				b.Fatalf("Has failed on key %s: %v", testKey, err)
			}
			if !has {
				b.Fatalf("Has reported key %s not present", testKey)
			}
		}
	}
}

func BenchmarkSimpleGet3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := []string{low, med, high}
		for _, key := range keys {
			testKey := []byte(key)
			_, err := simpleDB.Get(testKey)
			if err != nil {
				b.Fatalf("Failed to get key %s: %v", testKey, err)
			}
		}
	}
}

func BenchmarkSkipRangeScan(b *testing.B) {
	startKey := []byte("key0900")
	endKey := []byte("key0950")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator, err := skipDB.RangeScan(startKey, endKey)
		if err != nil {
			b.Fatalf("RangeScan failed: %v", err)
		}

		// Assuming iterator needs to be iterated to perform the scan.
		iters := 0
		for iterator.Next() {
			iters++
		}
		if i == b.N-1 {
			println("SkipList Iterations:", iters)
		}

		if err := iterator.Error(); err != nil {
			b.Fatalf("Iteration failed: %v", err)
		}
	}
}

func BenchmarkSimpleRangeScan(b *testing.B) {
	startKey := []byte("key0900")
	endKey := []byte("key0950")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator, err := simpleDB.RangeScan(startKey, endKey)
		if err != nil {
			b.Fatalf("RangeScan failed: %v", err)
		}

		// Iterate over the results.
		iters := 0
		for iterator.Next() {
			iters++
		}
		if i == b.N-1 {
			println("SimpleList Iterations:", iters)
		}

		if err := iterator.Error(); err != nil {
			b.Fatalf("Iteration failed: %v", err)
		}
	}
}

func BenchmarkSkipPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := []byte("key" + strconv.Itoa(i))
		value := []byte("value" + strconv.Itoa(i))
		skipDB.Put(key, value)
	}
}

func BenchmarkSimplePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := []byte("key" + strconv.Itoa(i))
		value := []byte("value" + strconv.Itoa(i))
		simpleDB.Put(key, value)
	}
}
