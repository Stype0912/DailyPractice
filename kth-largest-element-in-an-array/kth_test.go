package kth_largest_element_in_an_array

import "testing"

func TestKth(t *testing.T) {
	t.Log(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
