package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	return dfs(t1, t2)
}

func dfs(t1, t2 *TreeNode) bool {
	var boolean bool
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 == nil {
		return false
	}
	if t1 == nil && t2 == nil {
		return true
	}
	if equalTree(t1, t2) && t1.Val == t2.Val {
		return true
	}
	boolean = dfs(t1.Left, t2)
	if boolean {
		return true
	}
	boolean = dfs(t1.Right, t2)
	if boolean {
		return true
	}
	return false
}

func equalTree(root1, root2 *TreeNode) bool {
	var boolean bool
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1 == nil && root2 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	boolean = equalTree(root1.Left, root2.Left)
	if !boolean {
		return false
	}
	boolean = equalTree(root1.Right, root2.Right)
	if !boolean {
		return false
	}
	return true
}

func main() {
	fmt.Println(equalTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}))
}
