package evaluate_boolean_binary_tree

import "testing"

func TestTree(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3, Left: node2, Right: node3}
	node5 := &TreeNode{Val: 2, Left: node1, Right: node4}
	t.Log(evaluateTree(node5))
}
