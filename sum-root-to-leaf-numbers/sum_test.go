package sum_root_to_leaf_numbers

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Log(sumNumbers(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
}
