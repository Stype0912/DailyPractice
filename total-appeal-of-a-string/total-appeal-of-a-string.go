package total_appeal_of_a_string

func appealSum(s string) int64 {
	dp := make([]int, len(s)+1)
	lastPosMap := map[int32]int{}
	for i, letter := range s {
		j := -1
		if _, ok := lastPosMap[letter]; ok {
			j = lastPosMap[letter]
		}
		dp[i+1] = dp[i] + i - j
		lastPosMap[letter] = i
	}
	ans := 0
	for i := 0; i <= len(dp); i++ {
		ans += dp[i]
	}
	return int64(ans)
}
