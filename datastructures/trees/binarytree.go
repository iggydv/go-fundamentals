package trees

type BinaryTreeNode struct {
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (n *BinaryTreeNode) InsertLevelOrder(val int) {
	if n == nil {
		return
	}
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
