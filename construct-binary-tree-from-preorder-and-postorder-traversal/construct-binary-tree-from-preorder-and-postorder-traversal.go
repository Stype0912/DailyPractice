package construct_binary_tree_from_preorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var build func(int, int, int, int) *TreeNode
	idx := map[int]int{}
	for i, num := range postorder {
		idx[num] = i
	}
	build = func(preorderLeft, preorderRight, postorderLeft, postorderRight int) *TreeNode {
		if preorderLeft > preorderRight {
			return nil
		}
		leftCount := 0
		if preorderLeft < preorderRight {
			leftCount = idx[preorder[preorderLeft+1]] - postorderLeft + 1
		}
		return &TreeNode{
			Val:   preorder[preorderLeft],
			Left:  build(preorderLeft+1, preorderLeft+leftCount, postorderLeft, postorderLeft+leftCount-1),
			Right: build(preorderLeft+leftCount+1, preorderRight, postorderLeft+leftCount, postorderRight-1)}
	}
	return build(0, len(preorder)-1, 0, len(postorder)-1)
}
