package n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return ans
}
