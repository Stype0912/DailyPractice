package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		count := k
		stack := []*ListNode{}
		tmp := head
		for count > 0 && tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Next
			count--
		}
		if count > 0 {
			pre.Next = head
			break
		}
		for len(stack) != 0 {
			pre.Next = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pre = pre.Next
		}
		pre.Next = tmp
		head = tmp
		if tmp == nil {
			break
		}
	}
	return dummy.Next
}
