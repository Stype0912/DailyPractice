package maximum_subarray_sum_with_one_deletion

func maximumSum(arr []int) int {
	a, b, ans := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		b = max(b+arr[i], a)
		a = max(a+arr[i], arr[i])
		ans = max(max(a, ans), b)
	}
	return ans
}
