package maximum_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	dp := make([]int, len(nums))
	maxLeft := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[0]
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ans := nums[0]
	maxLeft[0] = nums[0]
	leftSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
		leftSum += nums[i]
		maxLeft[i] = max(maxLeft[i-1], leftSum)
	}
	rightSum := 0
	for i := len(nums) - 1; i >= 1; i-- {
		rightSum += nums[i]
		ans = max(ans, maxLeft[i-1]+rightSum)
	}
	return ans
}
