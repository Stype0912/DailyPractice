package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	var dfs func(node *TreeNode, target int, path []int)
	dfs = func(node *TreeNode, target int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Val == target && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int(nil), path...))
		}
		dfs(node.Left, target-node.Val, append([]int(nil), path...))
		dfs(node.Right, target-node.Val, append([]int(nil), path...))
	}
	dfs(root, targetSum, []int{})
	return ans
}
