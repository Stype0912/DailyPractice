package house_robber_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	f := map[*TreeNode]int{}
	g := map[*TreeNode]int{}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		f[node] = node.Val + g[node.Left] + g[node.Right]
		g[node] = max(f[node.Left], g[node.Left]) + max(f[node.Right], g[node.Right])
	}
	dfs(root)
	return max(f[root], g[root])
}
