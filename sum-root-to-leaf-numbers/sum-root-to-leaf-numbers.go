package sum_root_to_leaf_numbers

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, ancestors []int)
	ans := 0
	dfs = func(node *TreeNode, ancestors []int) {
		if node == nil {
			return
		}
		dfs(node.Left, append(ancestors, node.Val))
		dfs(node.Right, append(ancestors, node.Val))
		if node.Left == nil && node.Right == nil {
			newAncestors := append(ancestors, node.Val)
			ancestorStr := ""
			for _, ancestor := range newAncestors {
				ancestorStr += strconv.FormatInt(int64(ancestor), 10)
			}
			ancestorInt, _ := strconv.ParseInt(ancestorStr, 10, 64)
			ans += int(ancestorInt)
		}
	}
	dfs(root, []int{})
	return ans
}
