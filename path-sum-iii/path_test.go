package path_sum_iii

import "testing"

func TestPath(t *testing.T) {
	t.Log(pathSum(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Right: &TreeNode{Val: 11},
		},
	}, 8))
}
