package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	ans := []int{}
	ansRoot := &TreeNode{}
	tmp := ansRoot
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 0; i < len(ans); i++ {
		tmp.Right = &TreeNode{Val: ans[i]}
		tmp = tmp.Right
	}
	//*root = *ansRoot.Right
	root = ansRoot.Right
}
