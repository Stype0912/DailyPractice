package number_of_islands_ii

import "testing"

func TestIsland(t *testing.T) {
	t.Log(numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}))
}
