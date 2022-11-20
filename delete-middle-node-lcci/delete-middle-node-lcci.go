package delete_middle_node_lcci

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for {
		if node.Next.Next == nil {
			node.Val = node.Next.Val
			node.Next = nil
			break
		} else {
			node.Val = node.Next.Val
			node.Next.Val = node.Next.Next.Val
		}
		node = node.Next
	}
	node.Next = nil
}
