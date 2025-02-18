package main

import "testing"

func TestCollectionIterator_HasNext(t *testing.T) {
	collection := &Collection{
		items: []string{"a", "b", "c"},
	}
	it := &CollectionIterator{
		collection: collection,
		index:      0,
	}
	if !it.HasNext() {
		t.Errorf("Expected true but got false")
	}
	it.index = 3
	if it.HasNext() {
		t.Errorf("Expected false but got true")
	}
}

func TestCollectionIterator_Next(t *testing.T) {
	collection := &Collection{
		items: []string{"a", "b", "c"},
	}
	it := &CollectionIterator{
		collection: collection,
		index:      0,
	}
	item := it.Next()
	if item != "a" {
		t.Errorf("Expected a but got %s", item)
	}
	item = it.Next()
	if item != "b" {
		t.Errorf("Expected b but got %s", item)
	}
	item = it.Next()
	if item != "c" {
		t.Errorf("Expected c but got %s", item)
	}
	item = it.Next()
	if item != nil {
		t.Errorf("Expected nil but got %s", item)
	}
}
