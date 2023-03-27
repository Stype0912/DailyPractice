package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := [][]*TreeNode{{root}}
	ans := [][]int{}
	for len(stack[0]) != 0 {
		nodeOfFirstLevel := stack[0]
		newNodes := []*TreeNode{}
		newAns := []int{}
		for len(nodeOfFirstLevel) != 0 {
			node := nodeOfFirstLevel[0]
			newAns = append(newAns, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
			nodeOfFirstLevel = nodeOfFirstLevel[1:]
		}
		stack = append(stack, newNodes)
		ans = append(ans, newAns)
		stack = stack[1:]
	}
	return ans
}
