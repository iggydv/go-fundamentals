package main

import "testing"

func TestHashtable_Init(t *testing.T) {
	ht := &Hashtable{}
	ht.Init(10)
	if len(ht.array) != 10 {
		t.Errorf("Expected hashtable size of 10, got %d", len(ht.array))
	}
	for i, bucket := range ht.array {
		if bucket == nil {
			t.Errorf("Expected bucket at index %d to be initialized, got nil", i)
		}
	}
}

func TestHashtable_Insert(t *testing.T) {
	ht := &Hashtable{}
	ht.Init(10)
	ht.Insert("testKey")
	index := hash("testKey", len(ht.array))
	if !ht.array[index].search("testKey") {
		t.Errorf("Expected to find 'testKey' in hashtable")
	}
}

func TestHashtable_Search(t *testing.T) {
	ht := &Hashtable{}
	ht.Init(10)
	ht.Insert("testKey")
	if !ht.Search("testKey") {
		t.Errorf("Expected to find 'testKey' in hashtable")
	}
	if ht.Search("nonExistentKey") {
		t.Errorf("Did not expect to find 'nonExistentKey' in hashtable")
	}
}

func TestHashtable_Delete(t *testing.T) {
	ht := &Hashtable{}
	ht.Init(10)
	ht.Insert("testKey")
	ht.Delete("testKey")
	if ht.Search("testKey") {
		t.Errorf("Did not expect to find 'testKey' in hashtable after deletion")
	}
}
