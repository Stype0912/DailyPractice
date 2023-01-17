package add_two_numbers_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stackL1, stackL2 []int
	for {
		stackL1 = append(stackL1, l1.Val)
		if l1.Next == nil {
			break
		} else {
			l1 = l1.Next
		}
	}
	for {
		stackL2 = append(stackL2, l2.Val)
		if l2.Next == nil {
			break
		} else {
			l2 = l2.Next
		}
	}
	i := len(stackL1) - 1
	j := len(stackL2) - 1
	var jinwei = 0
	head := &ListNode{}
	res := head
	for {
		var elementI, elementJ int
		if i < 0 {
			elementI = 0
		} else {
			elementI = stackL1[i]
		}
		if j < 0 {
			elementJ = 0
		} else {
			elementJ = stackL2[j]
		}
		if i < 0 && j < 0 && jinwei == 0 {
			break
		}
		tmpNode := &ListNode{
			Val:  (elementJ + elementI + jinwei) % 10,
			Next: nil,
		}
		if res.Next == nil {
			res.Next = tmpNode
		} else {
			tmpNode.Next = res.Next
			res.Next = tmpNode
		}
		if elementJ+elementI+jinwei >= 10 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		i--
		j--
	}
	return res.Next
}
