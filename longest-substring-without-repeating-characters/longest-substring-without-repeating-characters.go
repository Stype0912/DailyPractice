package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	if n == 0 {
		return ans
	}
	start, end := 0, 0
	wordMap := map[byte]int{}
	for end < n {
		letter := s[end]
		if wordMap[letter] == 0 {
			ans = max(ans, end-start+1)

		} else {
			for wordMap[letter] != 0 {
				wordMap[s[start]]--
				start++
			}
		}
		end++
		wordMap[letter]++
	}
	return ans
}
