package queues

// Queue represents a queue data structure - FIFO (First In First Out)
// Time complexity:
// - Enqueue: O(1)
// - Dequeue: O(1)
// - Peek: O(1)
// - Search: O(n)
// - Remove: O(n)
// Space complexity: O(n)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}
