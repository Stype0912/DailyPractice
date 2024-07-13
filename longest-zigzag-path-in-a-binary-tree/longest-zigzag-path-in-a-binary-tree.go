package longest_zigzag_path_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, dir, steps int)
	dfs = func(node *TreeNode, dir int, steps int) {
		if node == nil {
			ans = max(ans, steps-2)
			return
		}
		if dir == 0 {
			dfs(node.Left, 1, steps+1)
			dfs(node.Right, 0, 2)
		} else {
			dfs(node.Right, 0, steps+1)
			dfs(node.Left, 1, 2)
		}
	}
	dfs(root, 0, 1)
	dfs(root, 1, 1)
	return ans
}
