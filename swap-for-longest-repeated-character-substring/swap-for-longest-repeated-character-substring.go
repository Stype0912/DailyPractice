package swap_for_longest_repeated_character_substring

func maxRepOpt1(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}

	res := 0
	i := 0
	for i < len(text) {
		j := i
		for j < len(text) && text[j] == text[i] {
			j++
		}
		curCnt := j - i
		if curCnt < count[rune(text[i])] && (j < len(text) || i > 0) {
			res = max(res, curCnt+1)
		}
		k := j + 1
		for k < len(text) && text[k] == text[i] {
			k++
		}
		res = max(res, min(k-i, count[rune(text[i])]))
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
