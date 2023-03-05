package shan_chu_lian_biao_de_jie_dian_lcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	nullNode := &ListNode{
		Next: head,
	}
	for {
		if nullNode.Next.Val == val {
			nullNode.Next = nullNode.Next.Next
			return head
		} else {
			nullNode = nullNode.Next
		}
	}
}
