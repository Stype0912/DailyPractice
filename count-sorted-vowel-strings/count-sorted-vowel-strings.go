package count_sorted_vowel_strings

func countVowelStrings(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[n-1][i]
	}
	return ans
}
