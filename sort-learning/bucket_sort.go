package sort_learning

func sortInBuckets(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func bucketSort(arr []int, min, max int) {
	defaultBucketSize := 2
	buckets := make([][]int, (max-min)/defaultBucketSize+1)
	for i := 0; i < len(arr); i++ {
		index := (arr[i] - min) / defaultBucketSize
		buckets[index] = append(buckets[index], arr[i])
	}
	pos := 0
	for i := 0; i < len(buckets); i++ {
		sortInBuckets(buckets[i])
		copy(arr[pos:], buckets[i])
		pos += len(buckets[i])
	}
	return
}
