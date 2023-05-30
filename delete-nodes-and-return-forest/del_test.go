package delete_nodes_and_return_forest

import "testing"

func TestDelete(t *testing.T) {
	t.Log(delNodes(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}, []int{3, 5}))
}
