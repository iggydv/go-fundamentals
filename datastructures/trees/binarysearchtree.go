package trees

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n.Key < val {
		if n.Right == nil {
			n.Right = &Node{Key: val}
		} else {
			n.Right.Insert(val)
		}
	} else if n.Key > val {
		if n.Left == nil {
			n.Left = &Node{Key: val}
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.Key < val {
		return n.Right.Search(val)
	}
	if n.Key > val {
		return n.Left.Search(val)
	}
	return true
}
