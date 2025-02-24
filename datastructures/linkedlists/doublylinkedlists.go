package linkedlists

// Time complexity:
// - Access: O(n)
// - Search: O(n)
// - Insertion: O(1)
// - Deletion: O(1)

// Space complexity: O(n)
// Doubly linked lists are a linear data structure that stores elements in a non-contiguous manner.
// They are useful when you need to store elements in a non-contiguous manner and when you need to perform frequent insertions and deletions.
// Since they require O(n) time to access an element, they are not suitable for applications that require frequent access to elements.

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}

type DoublyLinkedList struct {
	head   *DoubleNode
	tail   *DoubleNode
	length int
}

func (dl *DoublyLinkedList) Prepend(value int) {
	currentHead := dl.head
	n := &DoubleNode{
		next:  currentHead,
		prev:  nil,
		value: value,
	}

	if currentHead != nil {
		currentHead.prev = n
	} else {
		dl.tail = n
	}
	dl.head = n
	dl.length++
}

func (dl *DoublyLinkedList) Append(value int) {
	currentTail := dl.tail
	n := &DoubleNode{
		value: value,
		next:  nil,
		prev:  currentTail,
	}
	if currentTail != nil {
		currentTail.next = n
	} else {
		dl.head = n
	}
	dl.tail = n
	dl.length++
}

func (dl *DoublyLinkedList) Delete(value int) {
	if dl.head == nil {
		return
	}

	if dl.head.value == value {
		dl.head = dl.head.next
		dl.head.next.prev = nil
		dl.length--
		return
	}

	current := dl.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			if current.next != nil {
				current.next.prev = current
			} else {
				dl.tail = current
			}
			dl.length--
			return
		}
		current = current.next
	}
}
