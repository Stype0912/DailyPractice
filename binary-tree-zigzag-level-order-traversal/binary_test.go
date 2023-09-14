package binary_tree_zigzag_level_order_traversal

import "testing"

func TestBinary(t *testing.T) {
	t.Log(zigzagLevelOrder(&TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))
}
