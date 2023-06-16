package remove_zero_sum_consecutive_nodes_from_linked_list

import "testing"

func TestList(t *testing.T) {
	t.Log(removeZeroSumSublists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: -3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}))
}
