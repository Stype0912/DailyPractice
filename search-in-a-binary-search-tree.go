package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}
	return nil
}
