package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	node1, node2 := &ListNode{
		Next: head,
	}, &ListNode{
		Next: head,
	}
	for {
		node1 = node1.Next
		node2 = node2.Next
		if node2 == nil {
			return node1
		} else {
			node2 = node2.Next
		}
		if node2 == nil {
			return node1
		}
	}
}
