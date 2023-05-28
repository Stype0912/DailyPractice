package difference_of_number_of_distinct_values_on_diagonals

import "testing"

func TestDistinct(t *testing.T) {
	t.Log(differenceOfDistinctValues([][]int{
		{6, 28, 37, 34, 12, 30, 43, 35, 6},
		{21, 47, 38, 14, 31, 49, 11, 14, 49},
		{6, 12, 35, 17, 17, 2, 45, 27, 43},
		{34, 41, 30, 28, 45, 24, 50, 20, 4}}))
}
