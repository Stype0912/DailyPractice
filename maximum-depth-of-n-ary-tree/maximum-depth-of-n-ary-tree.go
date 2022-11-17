package main

type Node struct {
	Val      int
	Children []*Node
}

var ans = 0

func maxDepth(root *Node) int {
	dfsNNode(root, 0)
	return ans
}

func dfsNNode(root *Node, depth int) {
	if root == nil {
		return
	}
	depth += 1
	if depth > ans {
		ans = depth
	}
	for i := 0; i < len(root.Children); i++ {
		dfsNNode(root.Children[i], depth)
	}
}
