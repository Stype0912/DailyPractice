package copy_list_with_random_pointer

import "testing"

func TestRandom(t *testing.T) {
	node0 := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0
	t.Log(copyRandomList(node0))
}
