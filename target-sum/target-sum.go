package target_sum

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	neg := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (sum-target)%2 != 0 || sum-target < 0 {
		return 0
	} else {
		neg = (sum - target) / 2
	}
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for j := 1; j <= neg; j++ {
		dp[0][j] = 0
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= neg; j++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] + dp[i][j-nums[i]]
			}
		}
	}
	return dp[len(nums)][neg]
}
