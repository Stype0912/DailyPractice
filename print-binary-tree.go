package main

import (
	"math"
	"strconv"
)

var theMaxDepth = 0
var Map map[int][]int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	theMaxDepth = 0
	Map = make(map[int][]int)
	dfsDepth(root, 0)
	ans := make([][]string, theMaxDepth)
	for i, _ := range ans {
		ans[i] = make([]string, int64(math.Pow(float64(2), float64(theMaxDepth)))-1)
	}
	for i := theMaxDepth - 1; i >= 0; i-- {
		j := int64(math.Pow(float64(2), float64(theMaxDepth-1-i))) - 1
		k := 0
		for {
			if Map[i+1][k] == -9999 {
				ans[i][j] = ""
			} else {
				ans[i][j] = strconv.FormatInt(int64(Map[i+1][k]), 10)
			}
			k++
			j += int64(math.Pow(float64(2), float64(theMaxDepth-i)))
			if j >= int64(math.Pow(float64(2), float64(theMaxDepth))) || k == len(Map[i+1]) {
				break
			}
		}
	}
	return ans
}

func dfsDepth(root *TreeNode, depth int) {

	if root == nil {
		return
	}
	depth++
	if depth > theMaxDepth {
		theMaxDepth = depth
	}
	dfsDepth(root.Left, depth)
	dfsDepth(root.Right, depth)
}
