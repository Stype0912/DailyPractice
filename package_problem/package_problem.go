package package_problem

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func FindBestPackage(weight, value []int, content int) int {
	dp := make([][]int, len(weight)+1)
	for i := 0; i <= len(weight); i++ {
		dp[i] = make([]int, content+1)
	}
	for j := 0; j <= content; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= len(weight); i++ {
		for j := content; j >= 0; j-- {
			if j >= weight[i-1] {
				//fmt.Println(i, j)
				//fmt.Println(dp[i-1][j-weight[i-1]])
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(weight)][content]
}
