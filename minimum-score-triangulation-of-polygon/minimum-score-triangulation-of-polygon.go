package minimum_score_triangulation_of_polygon

import "math"

func minScoreTriangulation(values []int) int {
	var dp func(int, int) int
	memo := map[int]int{}
	dp = func(i int, j int) int {
		if i+2 > j {
			return 0
		}
		if i+2 == j {
			return values[i] * values[i+1] * values[i+2]
		}
		key := i*len(values) + j
		if _, ok := memo[key]; !ok {
			ans := math.MaxInt
			for k := i + 1; k < j; k++ {
				ans = min(ans, dp(i, k)+dp(k, j)+values[i]*values[j]*values[k])
			}
			memo[key] = ans
		}

		return memo[key]
	}
	return dp(0, len(values)-1)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
