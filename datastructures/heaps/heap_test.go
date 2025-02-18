package heaps

import (
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	h := &MaxHeap{}
	h.Insert(1)
	if h.items[0] != 1 {
		t.Errorf("Expected 1 but got %d", h.items[0])
	}
	h.Insert(2)
	if h.items[0] != 2 {
		t.Errorf("Expected 2 but got %d", h.items[0])
	}
	h.Insert(3)
	if h.items[0] != 3 {
		t.Errorf("Expected 3 but got %d", h.items[0])
	}
}

func TestMaxHeap_Extract(t *testing.T) {
	h := &MaxHeap{}
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	val := h.Extract()
	if val != 3 {
		t.Errorf("Expected 3 but got %d", val)
	}
	val = h.Extract()
	if val != 2 {
		t.Errorf("Expected 2 but got %d", val)
	}
	val = h.Extract()
	if val != 1 {
		t.Errorf("Expected 1 but got %d", val)
	}
	val = h.Extract()
	if val != -1 {
		t.Errorf("Expected -1 but got %d", val)
	}
}
