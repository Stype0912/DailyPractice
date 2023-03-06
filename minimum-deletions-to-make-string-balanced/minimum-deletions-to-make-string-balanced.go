package minimum_deletions_to_make_string_balanced

func minimumDeletions(s string) int {
	dp := make([]int, len(s))
	dp[0] = 0
	bCount := func(s0 uint8) int {
		if s0 == 'a' {
			return 0
		} else {
			return 1
		}
	}(s[0])
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i < len(s); i++ {
		if s[i] == 'b' {
			dp[i] = dp[i-1]
			bCount++
		} else {
			dp[i] = min(dp[i-1]+1, bCount)
		}
	}
	return dp[len(s)-1]
}
