package li_wu_de_zui_da_jie_zhi_lcof

func maxValue(grid [][]int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	row := len(grid)
	line := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, line)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < line; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < line; j++ {
			leftGift := dp[i-1][j]
			upGift := dp[i][j-1]
			dp[i][j] = max(leftGift, upGift) + grid[i][j]
		}
	}
	return dp[row-1][line-1]
}
