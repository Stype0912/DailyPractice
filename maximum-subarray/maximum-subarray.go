package maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[0]
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	maxSum := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}
