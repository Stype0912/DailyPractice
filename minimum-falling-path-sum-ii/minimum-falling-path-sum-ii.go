package minimum_falling_path_sum_ii

import "math"

func minFallingPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	row := len(grid)
	line := len(grid[0])
	for i := 0; i < row; i++ {
		dp[i] = make([]int, line)
	}
	for i := 0; i < line; i++ {
		dp[0][i] = grid[0][i]
	}
	for i := 1; i < row; i++ {
		for j := 0; j < line; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 1; i < row; i++ {
		for j := 0; j < line; j++ {
			for k := 0; k < line; k++ {
				if k == j {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+grid[i][j])
			}
		}
	}
	ans := math.MaxInt
	for j := 0; j < line; j++ {
		ans = min(ans, dp[row-1][j])
	}
	return ans
}
