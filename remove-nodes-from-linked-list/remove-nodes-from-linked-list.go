package remove_nodes_from_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	node := dummy.Next
	stack := []*ListNode{node}
	for {
		node = node.Next
		if node == nil {
			break
		}
		for len(stack) > 0 && node.Val > stack[len(stack)-1].Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node)
	}
	node = dummy
	for len(stack) != 0 {
		node.Next = stack[0]
		stack = stack[1:]
		node = node.Next
	}
	return dummy.Next
}
