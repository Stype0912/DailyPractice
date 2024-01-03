package sort_learning

import "testing"

func TestSort(t *testing.T) {
	arr := []int{13, 5, 6, 4, 2, 3, 65, 7, 8, 4, 2}
	quickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
