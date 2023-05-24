package frog_position_after_t_second

import "testing"

func TestFrog(t *testing.T) {
	t.Log(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
}
