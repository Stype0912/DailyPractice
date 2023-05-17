package sort_learning

import "testing"

func TestBucketSort(t *testing.T) {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8, 9, 10, -1, -3, -99999, 1000000}
	bucketSort(arr, -99999, 1000000)
	t.Log(arr)
}
