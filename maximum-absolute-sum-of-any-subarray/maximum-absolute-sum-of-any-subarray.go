package maximum_absolute_sum_of_any_subarray

func maxAbsoluteSum(nums []int) int {
	dp := make([]int, len(nums))
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	minSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = min(dp[i-1]+nums[i], nums[i])
		minSum = min(minSum, dp[i])
	}
	return max(maxSum, abs(minSum))
}
