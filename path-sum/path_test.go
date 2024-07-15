package main

import "testing"

func TestName(t *testing.T) {
	t.Log(hasPathSum(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: nil,
	}, 1))
}
