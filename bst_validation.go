package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	intMax := math.MaxInt64
	intMin := math.MinInt64
	return isValidBSTNode(root, intMin, intMax)
}

func isValidBSTNode(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	if !isValidBSTNode(node.Left, min, node.Val) {
		return false
	}
	if !isValidBSTNode(node.Right, node.Val, max) {
		return false
	}
	return true
}
