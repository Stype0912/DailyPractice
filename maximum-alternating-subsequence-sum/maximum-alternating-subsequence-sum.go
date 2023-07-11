package maximum_alternating_subsequence_sum

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	g := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = max(f[i-1], g[i-1]+nums[i-1])
		g[i] = max(g[i-1], f[i-1]-nums[i-1])
	}
	return int64(max(f[n], g[n]))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
