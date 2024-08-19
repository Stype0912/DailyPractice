package find_point_in_most_intervals

import "testing"

func TestIntervals(t *testing.T) {
	t.Log(findPointInMostIntervals([][]int{{0, 5}, {1, 6}, {2, 7}, {7, 11}, {11, 13}}))
}
