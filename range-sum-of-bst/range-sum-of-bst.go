package range_sum_of_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(node *TreeNode)
	ans := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val >= low && node.Val <= high {
			ans += node.Val
		}
		if node.Val > high {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
