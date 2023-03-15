package coin_change

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 0
	for j := 1; j <= amount; j++ {
		dp[0][j] = amount + 1
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[i+1][j] = min(dp[i][j], dp[i+1][j-coins[i]]+1)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	if dp[len(coins)][amount] < amount+1 {
		return dp[len(coins)][amount]
	} else {
		return -1
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
