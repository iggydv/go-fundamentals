package main

import "fmt"

type Hashtable struct {
	array []*Bucket
}

type Bucket struct {
	Head *BucketNode
}

// Linked list

type BucketNode struct {
	Key  string
	Next *BucketNode
}

func (h *Hashtable) Insert(key string) {
	index := hash(key, len(h.array))
	h.array[index].insert(key)
}

func (h *Hashtable) Search(key string) bool {
	index := hash(key, len(h.array))
	return h.array[index].search(key)
}

func (h *Hashtable) Delete(key string) {
	index := hash(key, len(h.array))
	h.array[index].delete(key)
}

func hash(key string, arrayLength int) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % arrayLength
}

// insert in head position
func (b *Bucket) insert(key string) {
	n := &BucketNode{Key: key}
	current := b.Head
	n.Next = current
	b.Head = n
}

func (b *Bucket) search(key string) bool {
	if b.Head == nil {
		return false
	}
	current := b.Head
	for current != nil {
		if current.Key == key {
			return true
		}
		current = current.Next
	}
	return false
}

func (b *Bucket) delete(key string) {
	if b.Head == nil {
		return
	}
	if b.Head.Key == key {
		b.Head = b.Head.Next
		return
	}
	current := b.Head
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (h *Hashtable) Init(size int) {
	if h.array == nil {
		h.array = make([]*Bucket, size)
		h.initBuckets()
	} else {
		fmt.Println("Cannot initialize an already initialized hashtable")
	}
}

func (h *Hashtable) initBuckets() {
	for i := range h.array {
		h.array[i] = &Bucket{}
	}
}
