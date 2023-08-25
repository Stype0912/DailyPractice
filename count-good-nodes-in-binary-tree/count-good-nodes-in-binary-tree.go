package count_good_nodes_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	ans := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var dfs func(node *TreeNode, maxFatherNode int)
	dfs = func(node *TreeNode, maxFatherNode int) {
		if node == nil {
			return
		}
		if maxFatherNode <= node.Val {
			ans++
		}
		maxFatherNode = max(maxFatherNode, node.Val)
		dfs(node.Left, maxFatherNode)
		dfs(node.Right, maxFatherNode)
	}
	dfs(root, -10001)
	return ans
}
