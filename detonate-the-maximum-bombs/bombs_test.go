package detonate_the_maximum_bombs

import "testing"

func TestBomb(t *testing.T) {
	t.Log(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))
}
