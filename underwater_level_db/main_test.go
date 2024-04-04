package main

import (
	"strconv"
	"testing"
	"underwater/simpleList"
	"underwater/skipList"
)

const (
	size = 10000
	low  = "key1"
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
	for i := 0; i < size; i++ {
		key := []byte("key" + strconv.Itoa(i))
		value := []byte("value" + strconv.Itoa(i))
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
	startKey := []byte("key100")
	endKey := []byte("key1100") // Adjust the range according to your test needs

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator, err := skipDB.RangeScan(startKey, endKey)
		if err != nil {
			b.Fatalf("RangeScan failed: %v", err)
		}

		// Assuming iterator needs to be iterated to perform the scan.
		for iterator.Next() {
		}

		if err := iterator.Error(); err != nil {
			b.Fatalf("Iteration failed: %v", err)
		}
	}
}

func BenchmarkSimpleRangeScan(b *testing.B) {
	startKey := []byte("key100")
	endKey := []byte("key1100") // Keep the range consistent for fair comparison

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator, err := simpleDB.RangeScan(startKey, endKey)
		if err != nil {
			b.Fatalf("RangeScan failed: %v", err)
		}

		// Iterate over the results.
		for iterator.Next() {
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
