package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced(root.Left) && isBalanced(root.Right) && (height(root.Left)-height(root.Right) >= -1 && height(root.Left)-height(root.Right) <= 1) {
		return true
	} else {
		return false
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
