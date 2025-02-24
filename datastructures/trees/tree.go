package trees

type TreeNode struct {
	Key      int
	Children []*TreeNode
}

// Can have any number of children

func (n *TreeNode) AddChild(value int) *TreeNode {
	child := &TreeNode{Key: value}
	n.Children = append(n.Children, child)
	return child
}

func (n *TreeNode) DFSPreOrder(val int) bool {
	if n == nil {
		return false
	}
	if n.Key == val {
		return true
	}
	for _, child := range n.Children {
		if child.DFSPreOrder(val) {
			return true
		}
	}
	return false
}

func (n *TreeNode) DFSPreOrderSlice() []int {
	var result []int
	if n == nil {
		return result
	}
	result = append(result, n.Key)
	for _, child := range n.Children {
		result = append(result, child.DFSPreOrderSlice()...)
	}
	return result
}

func (n *TreeNode) DFSPostOrder(val int) bool {
	if n == nil {
		return false
	}
	for _, child := range n.Children {
		if child.DFSPostOrder(val) {
			return true
		}
	}
	if n.Key == val {
		return true
	}
	return false
}

func (n *TreeNode) BFS(val int) bool {
	if n == nil {
		return false
	}
	queue := []*TreeNode{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Key == val {
			return true
		}
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
	return false
}

func (n *TreeNode) DFSPostOrderSlice() []int {
	if n == nil {
		return nil
	}
	var result []int
	for _, child := range n.Children {
		childResult := child.DFSPostOrderSlice()
		result = append(result, childResult...)
	}
	result = append(result, n.Key)
	return result
}
