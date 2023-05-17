package sort_learning

import "testing"

func TestSort(t *testing.T) {
	arr := []int{2, 4, 3}
	quickSort(arr, 0, 2)
	t.Log(arr)
}
