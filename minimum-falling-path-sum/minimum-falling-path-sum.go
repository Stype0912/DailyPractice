package minimum_falling_path_sum

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n)
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 0; i < n; i++ {
		dp[i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		newDp := make([]int, n)
		newDp[0] = min(dp[0], dp[1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			newDp[j] = min(min(dp[j-1], dp[j]), dp[j+1]) + matrix[i][j]
		}
		newDp[n-1] = min(dp[n-2], dp[n-1]) + matrix[i][n-1]
		copy(dp, newDp)
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dp[i])
	}
	return ans
}
