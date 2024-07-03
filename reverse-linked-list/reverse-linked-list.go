package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	ptr := head
	var prev, after *ListNode
	for ptr != nil {
		after = ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = after
	}
	return prev
}
