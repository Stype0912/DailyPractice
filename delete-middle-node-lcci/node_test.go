package delete_middle_node_lcci

import "testing"

func TestNode(t *testing.T) {
	Node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	Node3 := &ListNode{
		Val:  3,
		Next: Node4,
	}
	Node2 := &ListNode{
		Val:  2,
		Next: Node3,
	}
	Node1 := &ListNode{
		Val:  1,
		Next: Node2,
	}
	t.Log(Node1)
	deleteNode(Node2)
	t.Log(Node1.Next.Val)
}
