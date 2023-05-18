package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	wordMap := map[byte]int{}
	currLen := 0
	maxLen := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if wordMap[s[i]] == 0 {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			for wordMap[s[i]] > 0 {
				currLen--
				wordMap[s[j]]--
				j++
			}
			currLen++
		}
		wordMap[s[i]]++
	}
	return maxLen
}
