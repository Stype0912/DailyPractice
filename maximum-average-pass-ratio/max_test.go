package maximum_average_pass_ratio

import "testing"

func TestMax(t *testing.T) {
	t.Log(maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2))
}
