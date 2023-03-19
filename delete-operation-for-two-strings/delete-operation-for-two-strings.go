package delete_operation_for_two_strings

func minDistance(word1 string, word2 string) int {
	longestCommon := longestCommonSubsequence(word1, word2)
	return (len(word1) - longestCommon) + (len(word2) - longestCommon)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
