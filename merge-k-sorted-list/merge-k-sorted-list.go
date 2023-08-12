package merge_k_sorted_list

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := NodeHeap{}
	for _, node := range lists {
		if node != nil {
			h.Push(node)
		}
	}
	heap.Init(&h)
	ansNode := &ListNode{}
	ptr := ansNode
	for h.Len() != 0 {
		leastNode := heap.Pop(&h).(*ListNode)
		ptr.Next = &ListNode{Val: leastNode.Val}
		ptr = ptr.Next
		if leastNode.Next != nil {
			heap.Push(&h, leastNode.Next)
		}
	}
	return ansNode.Next
}
