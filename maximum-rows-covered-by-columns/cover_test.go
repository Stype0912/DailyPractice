package maximum_rows_covered_by_columns

import "testing"

func TestCover(t *testing.T) {
	t.Log(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
}
