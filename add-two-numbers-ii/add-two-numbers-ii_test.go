package add_two_numbers_ii

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	t.Log(addTwoNumbers(l1, l2).Val, addTwoNumbers(l1, l2).Next.Val, addTwoNumbers(l1, l2).Next.Next.Val)
}
