package shortest_subarray_to_be_removed_to_make_array_sorted

import "testing"

func TestSubarray(t *testing.T) {
	t.Log(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}
