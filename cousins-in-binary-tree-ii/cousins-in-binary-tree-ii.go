package cousins_in_binary_tree_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	brotherMap := map[*TreeNode]int{}
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		tmp := []*TreeNode{}
		for i := 0; i < n; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			if node.Left != nil && node.Right != nil {
				brotherMap[node.Left] = node.Right.Val
				brotherMap[node.Right] = node.Left.Val
			}
		}
		for i := 0; i < n; i++ {
			node := queue[i]
			node.Val = sum - node.Val - brotherMap[node]
		}
		queue = tmp
	}
	return root
}
