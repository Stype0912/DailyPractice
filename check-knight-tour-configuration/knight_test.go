package check_knight_tour_configuration

import "testing"

func TestKnight(t *testing.T) {
	t.Log(checkValidGrid([][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}))
}
