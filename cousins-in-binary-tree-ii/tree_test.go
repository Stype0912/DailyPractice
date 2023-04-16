package cousins_in_binary_tree_ii

import "testing"

func TestTree(t *testing.T) {
	t.Log(replaceValueInTree(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  9,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}
