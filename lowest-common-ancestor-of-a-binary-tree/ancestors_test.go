package lowest_common_ancestor_of_a_binary_tree

import "testing"

func TestAn(t *testing.T) {
	t.Log(lowestCommonAncestor(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}, &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}, &TreeNode{Val: 2,
		Left:  nil,
		Right: nil}))
}
