package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node1, node2 := head, head.Next
	for node2 != nil {
		for node2.Val == node1.Val {
			node2 = node2.Next
		}
		node1.Next = node2
		node1 = node2
		node2 = node2.Next
	}
	return head
}
