package number_of_provinces

import "testing"

func TestProvince(t *testing.T) {
	t.Log(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}
