package main

import (
	"fmt"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	// For Binary Search Trees, in-order traversal visits nodes in ascending order
	var values []int
	inorderTraversal(root, &values)

	fmt.Println(values)
	// Find minimum difference between adjacent values
	minDiff := math.MaxInt32
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func inorderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left, values)
	*values = append(*values, node.Val)
	inorderTraversal(node.Right, values)
}
