package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// Process nodes at current level
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// If this is the rightmost node at this level
			if i == levelSize-1 {
				result = append(result, current.Val)
			}

			// Add children to the queue (left first, then right)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return result
}
