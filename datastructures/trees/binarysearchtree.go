package trees

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if val > n.Val {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	} else if val < n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if val > n.Val {
		return n.Right.Search(val)
	}
	if val < n.Val {
		return n.Left.Search(val)
	}
	return true
}
