package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(node *TreeNode, lowerBound, upperBound int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lowerBound || node.Val >= upperBound {
		return false
	}
	return isValid(node.Left, lowerBound, node.Val) && isValid(node.Right, node.Val, upperBound)
}
