package main

var first, second, firstLevel, secondLevel, firstFather, secondFather int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	first = x
	second = y
	firstLevel, secondLevel, firstFather, secondFather = -1, -1, -1, -1
	bfs(root, 0, -1)
	if firstLevel == secondLevel && firstLevel != -1 && firstFather != secondFather {
		return true
	}
	return false
}

func bfs(root *TreeNode, level int, father int) {
	if root == nil {
		return
	}
	if root.Val == first {
		firstLevel = level
		firstFather = father
	}
	if root.Val == second {
		secondLevel = level
		secondFather = father
	}
	level += 1
	bfs(root.Left, level, root.Val)
	bfs(root.Right, level, root.Val)
}
