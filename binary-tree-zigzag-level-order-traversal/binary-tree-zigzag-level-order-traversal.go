package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) != 0 {
		queueLength := len(queue)
		tmp := []int{}
		for i := 0; i < queueLength; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i] = tmp[len(tmp)-i], tmp[i]
			}
		}
		ans = append(ans, tmp)
		level++
	}
	return ans
}
