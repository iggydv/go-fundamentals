package linkedlists

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
		return
	}
	return
}
