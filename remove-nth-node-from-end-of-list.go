package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	cur1, cur2 := dummy, dummy

	var counter int
	for {
		if counter >= n {
			cur1 = cur1.Next
		}
		cur2 = cur2.Next
		counter++
		if cur2.Next == nil {
			cur1.Next = cur1.Next.Next
			break
		}
	}
	return dummy.Next
}
