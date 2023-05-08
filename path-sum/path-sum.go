package main

var publicTargetSum int
var publicBool bool

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	publicTargetSum = targetSum
	publicBool = false
	dfsPathSum(root, 0)
	return publicBool
}

func dfsPathSum(root *TreeNode, val int) {
	if root == nil {
		return
	}
	temp := val + root.Val
	if root.Left == nil && root.Right == nil && temp == publicTargetSum {
		publicBool = true
	}
	dfsPathSum(root.Left, temp)
	dfsPathSum(root.Right, temp)
}
