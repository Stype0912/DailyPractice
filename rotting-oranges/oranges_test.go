package rotting_oranges

import "testing"

func TestOrange(t *testing.T) {
	t.Log(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}
