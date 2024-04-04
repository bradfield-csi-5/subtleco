package simpleList

import (
	"bytes"
	"testing"
)

// TestPutAndGet tests both the Put and Get methods of the Database.
func TestPutAndGet(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	key, value := []byte("Cat_01"), []byte("Parmesan")

	db.Put(key, value)
	val, err := db.Get(key)
	if err != nil {
		t.Fatalf("Get failed for key %s: %v", key, err)
	}
	if !bytes.Equal(val, value) {
		t.Errorf("Expected value '%s', got '%s'", value, val)
	}
}

// TestGetFail tests the Get method for a non-existent key.
func TestGetFail(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	_, err := db.Get([]byte("Nothing"))
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

// TestDelete tests the Delete method of the Database.
func TestDelete(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	key := []byte("Dog_01")
	db.Put(key, []byte("Frodo Waggins"))

	db.Delete(key)
	val, _ := db.Get(key)
	if len(val) > 0 {
		t.Errorf("Expected value to be blank, got '%s'", string(val))
	}
}

// TestHas tests the Has method for both existing and non-existing keys.
func TestHas(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	existingKey := []byte("Dog_02")
	nonExistingKey := []byte("Nothing")
	db.Put(existingKey, []byte("LC"))

	has, _ := db.Has(existingKey)
	if !has {
		t.Errorf("Expected to find key '%s', but it was not found", existingKey)
	}

	has, _ = db.Has(nonExistingKey)
	if has {
		t.Errorf("Did not expect to find key '%s', but it was found", nonExistingKey)
	}
}

// Test CSV seeding method
func TestCSVSeed(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	db.LoadCSV()

	key := "Grumpier Old Men (1995)"
	has, _ := db.Has([]byte(key))
	if !has {
		t.Errorf("Expected to find key '%s', but it was not found", key)
	}
}

// TestRangeScan tests the RangeScan method for a given range.
func TestRangeScan(t *testing.T) {
	db := Database{Entries: make([]Entry, 0)}
	db.Put([]byte("Cat_01"), []byte("Parmesan"))
	db.Put([]byte("Cat_02"), []byte("Tabasco"))
	db.Put([]byte("Cat_03"), []byte("Franklin"))
	db.Put([]byte("Cat_04"), []byte("Chef"))
	db.Put([]byte("Cat_05"), []byte("Garfield"))
	db.Put([]byte("Cat_06"), []byte("Idunno"))

	db.Put([]byte("Dog_01"), []byte("Frodo Waggins"))
	db.Put([]byte("Dog_02"), []byte("LC"))
	db.Put([]byte("Dog_03"), []byte("Ada"))
	db.Put([]byte("Dog_04"), []byte("Indiana Bones"))
	db.Put([]byte("Dog_05"), []byte("Grunglefunk"))
	db.Put([]byte("Dog_06"), []byte("AnotherDog"))

	iter, err := db.RangeScan([]byte("Cat_02"), []byte("Dog_03"))
	if err != nil {
		t.Fatalf("RangeScan failed: %v", err)
	}
	if string(iter.Key()) != "Cat_02" || string(iter.Value()) != "Tabasco" {
		t.Error("RangeScan did not return the expected first result 'Cat_02:Tabasco'")
	}
	if !iter.Next() || string(iter.Key()) != "Cat_03" || string(iter.Value()) != "Franklin" {
		t.Error("RangeScan did not return the expected second result 'Cat_03:Franklin'")
	}
	count := 0
	for iter.Next() {
		count++
	}
	if count != 5 {
		t.Errorf("Iterator had the wrong number of values left. Expected 5, got %d", count)
	}
}
