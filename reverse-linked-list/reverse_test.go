package reverse_linked_list

import "testing"

func TestReverse(t *testing.T) {
	t.Log(reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}))
}
