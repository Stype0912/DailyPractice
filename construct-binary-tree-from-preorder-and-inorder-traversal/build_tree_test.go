package construct_binary_tree_from_preorder_and_inorder_traversal

import "testing"

func TestBuildTree(t *testing.T) {
	t.Log(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
