package minimum_cost_tree_from_leaf_value

import "math"

func mctFromLeafValues(arr []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp, maxVal := make([][]int, len(arr)), make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr))
		maxVal[i] = make([]int, len(arr))
		maxVal[i][i] = arr[i]
	}
	for length := 2; length <= len(arr); length++ {
		for i := 0; i <= len(arr)-length; i++ {
			j := i + length - 1
			dp[i][j] = math.MaxInt
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], maxVal[i][k]*maxVal[k+1][j]+dp[i][k]+dp[k+1][j])
			}
			maxVal[i][j] = max(maxVal[i][j-1], arr[j])
		}
	}
	return dp[0][len(arr)-1]
}
