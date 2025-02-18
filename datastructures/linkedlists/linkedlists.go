package linkedlists

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(value int) {
	n := &Node{
		next:  l.head,
		value: value,
	}
	l.head = n
	l.length++
}

func (l *LinkedList) Append(value int) {
	n := &Node{
		next:  nil,
		value: value,
	}

	if l.head == nil {
		l.head = n
		l.length++
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	l.length++
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.length--
			return
		}
		current = current.next
	}
}
