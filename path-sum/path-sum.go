package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, target int)
	ans := false
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		if node.Val == target && node.Left == nil && node.Right == nil {
			ans = true
		}
		dfs(node.Left, target-node.Val)
		dfs(node.Right, target-node.Val)
	}
	dfs(root, targetSum)
	return ans
}
