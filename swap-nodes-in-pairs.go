package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	node := dummy
	for {
		if node.Next == nil || node.Next.Next == nil {
			return head
		}
		nodeTmp1, nodeTmp2 := new(ListNode), new(ListNode)
		nodeTmp1.Val = node.Next.Val
		nodeTmp2.Val = node.Next.Next.Val
		nodeTmp1.Next = node.Next.Next.Next
		nodeTmp2.Next = nodeTmp1
		node.Next = nodeTmp2
		node = nodeTmp1
	}
}
