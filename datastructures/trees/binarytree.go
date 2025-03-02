package trees

type BinaryTreeNode struct {
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (n *BinaryTreeNode) BFS(val int) bool {
	if n == nil {
		return false
	}
	queue := []*BinaryTreeNode{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Key == val {
			return true
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return false
}

func (n *BinaryTreeNode) DFSPreOrderSearch(val int) bool {
	if n == nil {
		return false
	}
	if n.Key == val {
		return true
	}
	if n.Left.DFSPreOrderSearch(val) {
		return true
	}
	if n.Right.DFSPreOrderSearch(val) {
		return true
	}
	return false
}

func (n *BinaryTreeNode) DFSPreOrderSlice() []int {
	var result []int
	if n == nil {
		return result
	}
	result = append(result, n.Key)
	result = append(result, n.Left.DFSPreOrderSlice()...)
	result = append(result, n.Right.DFSPreOrderSlice()...)
	return result
}

func (n *BinaryTreeNode) DFSPostOrderSearch(val int) bool {
	if n == nil {
		return false
	}
	if n.Left.DFSPostOrderSearch(val) {
		return true
	}
	if n.Right.DFSPostOrderSearch(val) {
		return true
	}
	if n.Key == val {
		return true
	}
	return false
}

func (n *BinaryTreeNode) DFSPostOrderSlice() []int {
	var result []int
	if n == nil {
		return result
	}
	result = append(result, n.Left.DFSPostOrderSlice()...)
	result = append(result, n.Right.DFSPostOrderSlice()...)
	result = append(result, n.Key)
	return result
}

func (n *BinaryTreeNode) InsertLevelOrder(val int) {
	queue := []*BinaryTreeNode{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTreeNode{Key: val}
			return
		}
		queue = append(queue, current.Left)
		if current.Right == nil {
			current.Right = &BinaryTreeNode{Key: val}
			return
		}
		queue = append(queue, current.Right)
	}
}
