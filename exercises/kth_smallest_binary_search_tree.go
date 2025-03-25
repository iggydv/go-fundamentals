package main

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var f func(n *TreeNode)
	var count int
	var result int
	var found bool

	f = func(n *TreeNode) {
		if n == nil || found == true {
			return
		}

		f(n.Left)
		count++
		if count == k {
			found = true
			result = n.Val
		}
		f(n.Right)
	}

	f(root)
	return result
}
