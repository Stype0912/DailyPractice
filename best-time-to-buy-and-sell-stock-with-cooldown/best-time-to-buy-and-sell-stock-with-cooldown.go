package best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	max := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}
