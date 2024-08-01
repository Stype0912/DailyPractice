package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{Val: head.Val}
	newHeadPtr := newHead
	headPtr := head
	nodeMap := map[*Node]*Node{headPtr: newHeadPtr}
	for headPtr.Next != nil {
		newHeadPtr.Next = &Node{Val: headPtr.Next.Val, Next: nil, Random: nil}
		newHeadPtr = newHeadPtr.Next
		headPtr = headPtr.Next
		nodeMap[headPtr] = newHeadPtr
	}
	headPtr = head
	newHeadPtr = newHead
	for headPtr != nil {
		newHeadPtr.Random = nodeMap[headPtr.Random]
		newHeadPtr = newHeadPtr.Next
		headPtr = headPtr.Next
	}
	return newHead
}
