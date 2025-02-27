package heaps

// Implementing a priority queue with a hashtable

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(val int) {
	h.items = append(h.items, val)
	h.heapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.items) == 0 {
		return -1
	}
	extracted := h.items[0]
	l := len(h.items) - 1

	h.items[0] = h.items[l]
	h.items = h.items[:l]
	h.heapifyDown(0)
	return extracted
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	last := len(h.items) - 1
	l, r := left(index), right(index)
	current := 0

	for l <= last {
		if l == last {
			current = l
		} else if h.items[l] > h.items[r] {
			current = l
		} else {
			current = r
		}
		// compare and swap
		if h.items[index] < h.items[current] {
			h.swap(index, current)
			index = current
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2
}

func (h *MaxHeap) swap(i int, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
