package minimum_path_sum

func minPathSum(grid [][]int) int {
	row := len(grid)
	line := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, line)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			var tmp1, tmp2 int = -1, -1
			if i > 0 {
				tmp1 = dp[i-1][j] + grid[i-1][j]
			}
			if j > 0 {
				tmp2 = dp[i][j-1] + grid[i][j-1]
			}
			if tmp1 != -1 && tmp2 != -1 {
				dp[i][j] = min(tmp1, tmp2)
			} else if tmp1 == -1 {
				dp[i][j] = tmp2
			} else if tmp2 == -1 {
				dp[i][j] = tmp1
			}
		}
	}
	return dp[row-1][line-1]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
