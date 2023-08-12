package merge_k_sorted_list

import "testing"

func TestMerge(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	t.Log(mergeKLists([]*ListNode{list1, list2}))
}
