package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var build func(int, int) *TreeNode
	idx := map[int]int{}
	for i, num := range inorder {
		idx[num] = i
	}
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		if len(postorder) <= 0 {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		node := &TreeNode{Val: val}
		index := idx[val]
		node.Right = build(index+1, inorderRight)
		node.Left = build(inorderLeft, index-1)
		return node
	}
	return build(0, len(inorder)-1)
}
