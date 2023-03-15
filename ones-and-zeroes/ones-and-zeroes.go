package ones_and_zeroes

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i <= len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 0; i < len(strs); i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				zero, one := countZeroAndOne(strs[i])
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zero && k >= one && dp[i][j-zero][k-one]+1 > dp[i][j][k] {
					dp[i+1][j][k] = dp[i][j-zero][k-one] + 1
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func countZeroAndOne(str string) (int, int) {
	counter := map[int32]int{}
	for _, num := range str {
		counter[num]++
	}
	return counter['0'], counter['1']
}
