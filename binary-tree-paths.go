package main

import "strconv"

var ansStr []string

func binaryTreePaths(root *TreeNode) []string {
	ansStr = []string{}
	dfsPath(root, "")
	return ansStr
}

func dfsPath(root *TreeNode, path string) {
	if root == nil {
		return
	}
	if path == "" {
		path += strconv.FormatInt(int64(root.Val), 10)
	}else {
		path += "->" + strconv.FormatInt(int64(root.Val), 10)
	}
	if root.Left == nil && root.Right == nil {
		ansStr = append(ansStr, path)
	}
	dfsPath(root.Left, path)
	dfsPath(root.Right, path)
}
