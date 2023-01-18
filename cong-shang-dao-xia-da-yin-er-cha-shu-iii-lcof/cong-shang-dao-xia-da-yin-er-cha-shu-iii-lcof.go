package cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

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

func levelOrder(root *TreeNode) (ans [][]int) {
	var nodeOfLevel []*TreeNode
	if root != nil {
		nodeOfLevel = append(nodeOfLevel, root)
	}
	level := 0
	for {
		if len(nodeOfLevel) == 0 {
			break
		}
		res := []int{}
		nodeOfLevelTmp := []*TreeNode{}
		for _, node := range nodeOfLevel {
			res = append(res, node.Val)
			if node.Left != nil {
				nodeOfLevelTmp = append(nodeOfLevelTmp, node.Left)
			}
			if node.Right != nil {
				nodeOfLevelTmp = append(nodeOfLevelTmp, node.Right)
			}
		}
		nodeOfLevel = []*TreeNode{}
		nodeOfLevel = append(nodeOfLevel, nodeOfLevelTmp...)
		if level%2 == 0 {
			ans = append(ans, res)
		} else {
			ans = append(ans, reverse(res))
		}
		level++
	}
	return
}

func reverse(res []int) []int {
	newRes := []int{}
	for i := 0; i < len(res); i++ {
		newRes = append(newRes, res[len(res)-1-i])
	}
	return newRes
}
