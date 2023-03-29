package populating_next_right_pointers_in_each_node

import "testing"

func TestNext(t *testing.T) {
	t.Log(connect(&Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  nil,
			Right: nil,
			Next:  nil,
		},
		Right: &Node{
			Val:   3,
			Left:  nil,
			Right: nil,
			Next:  nil,
		},
		Next: nil,
	}))
}
