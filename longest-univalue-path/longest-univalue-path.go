package longest_univalue_path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftLength := dfs(node.Left)
		rightLength := dfs(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			leftLength++
		} else {
			leftLength = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			rightLength++
		} else {
			rightLength = 0
		}
		if leftLength+rightLength > ans {
			ans = leftLength + rightLength
		}
		return max(leftLength, rightLength)
	}
	dfs(root)
	return ans
}
