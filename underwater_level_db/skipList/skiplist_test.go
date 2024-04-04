package skipList

import (
	"strconv"
	"testing"
)

func BenchmarkPut(b *testing.B) {
	db, _ := CreateDatabase()
	b.ResetTimer() // Start the timer here to only measure the Put operation.
	for i := 0; i < b.N; i++ {
		key := []byte("key" + strconv.Itoa(i))
		value := []byte("value" + strconv.Itoa(i))
		db.Put(key, value)
	}
}
