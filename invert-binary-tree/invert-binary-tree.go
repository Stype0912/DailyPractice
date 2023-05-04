package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil && root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	} else if root.Left == nil && root.Right != nil {
		root.Left = root.Right
		root.Right = nil
		invertTree(root.Left)
	} else if root.Left != nil && root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		invertTree(root.Right)
	}
	return root
}
