package insufficient_nodes_in_root_to_leaf_paths

import "testing"

func TestInsufficient(t *testing.T) {
	t.Log(sufficientSubset(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   -99,
				Left:  &TreeNode{Val: -99},
				Right: &TreeNode{Val: -99},
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   -99,
				Left:  &TreeNode{Val: 12},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: -99},
				Right: &TreeNode{Val: 14},
			},
		},
	}, 1))
	//t.Log(sufficientSubset(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{Val: -5},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   -3,
	//		Left:  &TreeNode{Val: 4},
	//		Right: nil,
	//	},
	//}, -1))
}
