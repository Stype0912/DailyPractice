package intersection_of_two_linked_lists_lcci

import "testing"

func TestList(t *testing.T) {
	headA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	t.Log(getIntersectionNode(headA, headA.Next))
}
