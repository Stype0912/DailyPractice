package reverse_odd_levels_of_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	level := 0
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) > 0 {
		n := len(nodeQueue)
		tmp := []*TreeNode{}
		for i := 0; i < n; i++ {
			node := nodeQueue[0]
			nodeQueue = nodeQueue[1:]
			tmp = append(tmp, node)
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		if level&1 == 1 {
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i].Val, tmp[len(tmp)-i-1].Val = tmp[len(tmp)-i-1].Val, tmp[i].Val
			}
		}
		level++
	}
	return root
}
