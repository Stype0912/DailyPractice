package evaluate_boolean_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		return valOfNode(root.Val)
	}
	switch root.Val {
	case 2:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case 3:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	default:
		return valOfNode(root.Val)
	}
}

func valOfNode(val int) bool {
	if val == 0 {
		return false
	} else if val == 1 {
		return true
	} else {
		return true
	}
}
