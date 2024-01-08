package delete_node_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
		return root
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
		return root
	case root.Left == nil || root.Right == nil:
		if root.Right != nil {
			return root.Right
		}
		return root.Left
	default:
		tmp := root.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		root.Val = tmp.Val
		root.Right = deleteNode(root.Right, tmp.Val)
		return root
	}
}
