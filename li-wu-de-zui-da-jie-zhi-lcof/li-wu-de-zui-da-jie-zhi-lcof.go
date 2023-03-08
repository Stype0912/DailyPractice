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
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			leftGift, upGift := 0, 0
			if i > 0 {
				leftGift = dp[i-1][j]
			}
			if j > 0 {
				upGift = dp[i][j-1]
			}
			dp[i][j] = max(leftGift, upGift) + grid[i][j]
		}
	}
	return dp[row-1][line-1]
}
