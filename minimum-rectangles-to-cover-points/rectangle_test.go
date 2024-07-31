package minimum_rectangles_to_cover_points

import "testing"

func TestRectangle(t *testing.T) {
	t.Log(minRectanglesToCoverPoints([][]int{{1, 3}, {7, 3}}, 1))
}
