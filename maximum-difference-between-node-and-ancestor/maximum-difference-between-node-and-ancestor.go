package maximum_difference_between_node_and_ancestor

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, fatherVals []int) {
		if node == nil {
			return
		}
		for _, fatherVal := range fatherVals {
			ans = max(ans, int(math.Abs(float64(fatherVal-node.Val))))
		}
		dfs(node.Left, append(fatherVals, node.Val))
		dfs(node.Right, append(fatherVals, node.Val))
	}
	dfs(root, []int{})
	return ans
}
