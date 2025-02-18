package main

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection struct {
	items []string
}

type CollectionIterator struct {
	collection *Collection
	index      int
}

// Implement the interface

func (it *CollectionIterator) HasNext() bool {
	return it.index < len(it.collection.items)
}

func (it *CollectionIterator) Next() interface{} {
	if it.HasNext() {
		item := it.collection.items[it.index]
		it.index++
		return item
	}
	return nil
}
