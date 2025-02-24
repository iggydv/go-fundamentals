package stacks

// Stack is a basic stack implementation - Last In, First Out (LIFO)
// Time complexity:
// - Push: O(1)
// - Pop: O(1)
// - Peek: O(1)
// - Search: O(n)
// - Remove: O(n)
// Space complexity: O(n)

// Stacks are useful when you need to store elements in a contiguous manner and when you need to perform frequent insertions and deletions.

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	l := len(s.items)
	if l == 0 {
		return -1
	}
	last := s.items[l-1]
	s.items = s.items[:l-1]
	return last
}
