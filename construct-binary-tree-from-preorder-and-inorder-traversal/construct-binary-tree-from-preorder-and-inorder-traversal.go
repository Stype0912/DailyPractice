package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(int, int) *TreeNode
	idx := map[int]int{}
	for i, num := range inorder {
		idx[num] = i
	}
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := preorder[0]
		preorder = preorder[1:]
		node := &TreeNode{Val: val}
		index := idx[val]
		node.Left = build(inorderLeft, index-1)
		node.Right = build(index+1, inorderRight)
		return node
	}
	return build(0, len(inorder)-1)
}
