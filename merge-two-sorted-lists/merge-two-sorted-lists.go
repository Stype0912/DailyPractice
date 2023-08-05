package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newNode := &ListNode{}
	res := newNode
	for list1 != nil && list2 != nil {
		if list1 == nil {
			newNode.Next = &ListNode{Val: list2.Val}
			newNode = newNode.Next
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			newNode.Next = &ListNode{Val: list1.Val}
			newNode = newNode.Next
			list1 = list1.Next
			continue
		}
		if list1.Val < list2.Val {
			newNode.Next = &ListNode{Val: list1.Val}
			newNode = newNode.Next
			list1 = list1.Next
		} else {
			newNode.Next = &ListNode{Val: list2.Val}
			newNode = newNode.Next
			list2 = list2.Next
		}
	}
	return res.Next
}
