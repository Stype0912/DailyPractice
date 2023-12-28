package collecting_chocolates

func minCost(nums []int, x int) int64 {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	sum := func(arr []int) (ans int) {
		for _, item := range arr {
			ans += item
		}
		return ans
	}
	n := len(nums)
	f := make([][]int, n)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
		f[0][i] = nums[i]
	}
	ans := sum(nums)
	for k := 1; k < n; k++ {
		for j := 0; j < n; j++ {
			f[k][j] = min(f[k-1][j], nums[(j+k)%n])
		}
		ans = min(ans, k*x+sum(f[k]))
	}
	return int64(ans)
}
