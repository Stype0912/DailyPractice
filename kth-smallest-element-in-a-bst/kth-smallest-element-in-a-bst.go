package kth_smallest_element_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode)
	ans := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		} else if k < 0 {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
