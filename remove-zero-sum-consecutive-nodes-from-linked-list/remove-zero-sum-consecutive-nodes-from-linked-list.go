package remove_zero_sum_consecutive_nodes_from_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	nodeMap := map[int][]*ListNode{}
	dummy := &ListNode{Next: head}
	node := dummy
	prefixSum := 0
	for node != nil {
		prefixSum += node.Val
		nodeMap[prefixSum] = append(nodeMap[prefixSum], node)
		node = node.Next
	}
	node = dummy
	prefixSum = 0
	for node != nil {
		prefixSum += node.Val
		if len(nodeMap[prefixSum]) > 1 {
			lastNode := nodeMap[prefixSum][len(nodeMap[prefixSum])-1]
			node.Next = lastNode.Next
		}
		node = node.Next
	}
	return dummy.Next
}
