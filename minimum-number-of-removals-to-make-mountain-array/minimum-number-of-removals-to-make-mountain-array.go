package minimum_number_of_removals_to_make_mountain_array

func minimumMountainRemovals(nums []int) int {
	prefix := lengthOfLIS(nums)
	suffix := lengthOfLIS(reverse(nums))
	suffix = reverse(suffix)
	ans := 0
	n := len(nums)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < n; i++ {
		if prefix[i] > 1 && suffix[i] > 1 {
			ans = max(ans, prefix[i]+suffix[i]-1)
		}
	}
	return n - ans
}
func lengthOfLIS(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}

func reverse(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[n-i-1]
	}
	return ans
}
