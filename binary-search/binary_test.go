package binary_search

import "testing"

func TestSearch(t *testing.T) {
	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
