package intersection_of_two_linked_lists_lcci

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		hashMap[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if hashMap[tmp] {
			return tmp
		}
	}
	return nil
}
