package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	values := []int{}
	ansNode := &ListNode{}
	ptr := ansNode
	for head != nil {
		if head.Val < x {
			tmpNode := &ListNode{Val: head.Val}
			ptr.Next = tmpNode
			ptr = ptr.Next
		} else {
			values = append(values, head.Val)
		}
		head = head.Next
	}
	for i := 0; i < len(values); i++ {
		tmpNode := &ListNode{Val: values[i]}
		ptr.Next = tmpNode
		ptr = ptr.Next
	}
	return ansNode.Next
}
