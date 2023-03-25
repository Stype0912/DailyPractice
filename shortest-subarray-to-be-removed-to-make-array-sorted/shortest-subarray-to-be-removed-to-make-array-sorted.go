package shortest_subarray_to_be_removed_to_make_array_sorted

func findLengthOfShortestSubarray(arr []int) int {
	left := 0
	right := len(arr) - 1
	for i := 0; i < len(arr)-1 && arr[i] <= arr[i+1]; i++ {
		left++
	}
	for j := right; j > 0 && arr[j] >= arr[j-1]; j-- {
		right--
	}
	if left >= right {
		return 0
	}
	ans := min(len(arr)-left-1, right)

	for i := 0; i <= left; i++ {
		k := right
		for k < len(arr) && arr[k] < arr[i] {
			k++
		}
		ans = min(ans, k-i-1)
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
