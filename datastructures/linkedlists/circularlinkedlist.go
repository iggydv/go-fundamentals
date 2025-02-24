package linkedlists

// Time complexity:
// - Access: O(n)
// - Search: O(n)
// - Insertion: O(1)
// - Deletion: O(1)

// Space complexity: O(n)
// Circular linked lists are a linear data structure that stores elements in a non-contiguous manner.
// They are useful when you need to store elements in a non-contiguous manner and when you need to perform frequent insertions and deletions.
// Since they require O(n) time to access an element, they are not suitable for applications that require frequent access to elements.

type CircularLinkedList struct {
	head   *Node
	length int
}

func (cl *CircularLinkedList) Prepend(value int) {
	n := &Node{
		value: value,
	}
	if cl.head == nil {
		n.next = n
	} else {
		current := cl.head
		for current.next != cl.head {
			current = current.next
		}
		n.next = cl.head
		current.next = n
		cl.head = n
	}
	cl.length++
}

func (cl *CircularLinkedList) Append(value int) {
	n := &Node{
		value: value,
	}
	if cl.head == nil {
		n.next = n
		cl.head = n
	} else {
		current := cl.head
		for current.next != cl.head {
			current = current.next
		}
		current.next = n
	}
	cl.length++
}

func (cl *CircularLinkedList) Delete(value int) {
	if cl.head == nil {
		return
	}

	if cl.head.value == value {
		if cl.head.next == cl.head {
			cl.head = nil
		} else {
			current := cl.head
			for current.next != cl.head {
				current = current.next
			}
			current.next = cl.head.next
			cl.head = cl.head.next
		}
		cl.length--
		return
	}

	current := cl.head
	for current.next != cl.head && current.next.value != value {
		current = current.next
	}

	if current.next != cl.head {
		current.next = current.next.next
		cl.length--
	}
	return
}
