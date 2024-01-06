package insert_greatest_common_divisors_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	maxPublic := func(i, j int) int {
		n := min(i, j)
		for k := n; k >= 1; k-- {
			if i%k == 0 && j%k == 0 {
				return k
			}
		}
		return 1
	}
	dummy := &ListNode{Next: head}
	prev := dummy.Next
	next := dummy.Next.Next
	for next != nil {
		insertedVal := &ListNode{Val: maxPublic(prev.Val, next.Val), Next: next}
		prev.Next = insertedVal
		prev = next
		next = next.Next
	}
	return dummy.Next
}
