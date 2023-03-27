package binary_tree_maximum_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ans := (-1) * int(math.Pow(2, 32))
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(dfs(node.Left), 0)
		rightSum := max(dfs(node.Right), 0)
		ans = max(ans, node.Val+leftSum+rightSum)
		return max(leftSum+node.Val, rightSum+node.Val)
	}
	dfs(root)
	return ans
}
