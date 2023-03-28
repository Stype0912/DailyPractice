package flatten_binary_tree_to_linked_list

import "testing"

func TestFlatten(t *testing.T) {
	node := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	flatten(node)
	t.Log(node.Val)
}
