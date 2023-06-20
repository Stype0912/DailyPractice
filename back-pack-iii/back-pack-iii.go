package back_pack_iii

func BackPackIII(a []int, v []int, m int) int {
	// write your code here
	row := len(a)
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = v[0] * j / a[0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j <= m; j++ {
			maxTmp := 0
			for k := 0; k < row; k++ {
				if j-a[k] < 0 {
					continue
				} else {
					//maxTmp = max(maxTmp, dp[i-1][j-a[k]]+v[k])
					maxTmp = max(maxTmp, dp[i][j-a[k]]+v[k])
				}
			}
			dp[i][j] = max(dp[i-1][j-1], maxTmp)
		}
	}
	return dp[row-1][m]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
