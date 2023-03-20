package combination_sum_iv

func combinationSum4(nums []int, target int) int {
	row := len(nums)
	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i <= row; i++ {
		dp[i][0] = 1
	}
	for j := 1; j <= target; j++ {
		for i := 0; i < row; i++ {
			//dp[i+1][j] = dp[i][j]
			for k := 0; k <= i; k++ {
				if j >= nums[k] {
					dp[i+1][j] += dp[i+1][j-nums[k]]
				}
			}
		}
	}
	return dp[row][target]
}
