package min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
