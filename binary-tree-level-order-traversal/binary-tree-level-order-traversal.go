package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		ansTmp := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			ansTmp = append(ansTmp, node.Val)
		}
		ans = append(ans, ansTmp)
	}
	return ans
}
