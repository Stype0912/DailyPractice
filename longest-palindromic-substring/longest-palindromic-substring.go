package longest_palindromic_substring

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen := 1
	maxStart := 0
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := i + L - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxStart = i
				maxLen = j - i + 1
			}
		}
	}
	return s[maxStart : maxStart+maxLen]
}
