package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	ans := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}
