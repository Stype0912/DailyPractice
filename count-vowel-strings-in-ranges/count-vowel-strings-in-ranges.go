package count_vowel_strings_in_ranges

func vowelStrings(words []string, queries [][]int) []int {
	preSum := map[int]int{}
	preSum[-1] = 0
	if isValid(words[0]) {
		preSum[0] = 1
	}
	for i := 1; i < len(words); i++ {
		if isValid(words[i]) {
			preSum[i] = preSum[i-1] + 1
		} else {
			preSum[i] = preSum[i-1]
		}
	}
	ans := []int{}
	for i := 0; i < len(queries); i++ {
		ans = append(ans, preSum[queries[i][1]]-preSum[queries[i][0]-1])
	}
	return ans
}

func isValid(word string) bool {
	last := len(word) - 1
	if (word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u') &&
		(word[last] == 'a' || word[last] == 'e' || word[last] == 'i' || word[last] == 'o' || word[last] == 'u') {
		return true
	}
	return false
}
