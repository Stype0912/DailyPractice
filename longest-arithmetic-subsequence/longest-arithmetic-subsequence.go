package longest_arithmetic_subsequence

func longestArithSeqLength(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 1001)
	}
	ans := 0
	for i := 1; i < len(nums); i++ {
		for k := 0; k < i; k++ {
			j := nums[i] - nums[k] + 500
			dp[i][j] = max(dp[i][j], dp[k][j]+1)
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
