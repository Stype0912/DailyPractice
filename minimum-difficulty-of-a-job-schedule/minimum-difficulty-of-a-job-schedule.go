package minimum_difficulty_of_a_job_schedule

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	dp := make([][]int, d)
	for i := 0; i < d; i++ {
		dp[i] = make([]int, len(jobDifficulty))
	}
	for i := 0; i < d; i++ {
		for j := 0; j < len(jobDifficulty); j++ {
			dp[i][j] = math.MaxInt
		}
	}
	maxDifficulty := jobDifficulty[0]
	preSum := make([]int, len(jobDifficulty))
	preSum[0] = jobDifficulty[0]
	dp[0][0] = maxDifficulty
	for j := 1; j < len(jobDifficulty); j++ {
		maxDifficulty = max(maxDifficulty, jobDifficulty[j])
		dp[0][j] = maxDifficulty
		preSum[j] = preSum[j-1] + jobDifficulty[j]
	}
	//maxMap := map[int]map[int]int{}
	//maxFromItoJ := func(i, j int) int {
	//	ans := 0
	//	if ans, ok := maxMap[i][j]; ok {
	//		return ans
	//	}
	//	for k := i; k <= j; k++ {
	//		ans = max(ans, jobDifficulty[k])
	//	}
	//	return ans
	//}
	for i := 0; i < d-1; i++ {
		for j := 0; j < len(jobDifficulty); j++ {
			if i+1 > j {
				dp[i+1][j] = -1
			} else {
				maxTmp := 0
				for k := j - 1; k >= i; k-- {
					maxTmp = max(maxTmp, jobDifficulty[k+1])
					dp[i+1][j] = min(dp[i+1][j], dp[i][k]+maxTmp)
				}
			}
		}
	}
	return dp[d-1][len(jobDifficulty)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
