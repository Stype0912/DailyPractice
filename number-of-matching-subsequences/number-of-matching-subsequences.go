package number_of_matching_subsequences

import "sort"

func numMatchingSubseq(s string, words []string) int {
	var pos = [26][]int{}
	for i, letter := range s {
		pos[letter-'a'] = append(pos[letter-'a'], i)
	}
	ans := len(words)
	for _, word := range words {
		if len(word) > len(s) {
			ans--
		}
		ptr := -1
		for _, letter := range word {
			j := sort.SearchInts(pos[letter-'a'], ptr+1)
			if j == len(pos[letter-'a']) {
				ans--
				break
			}
			ptr = pos[letter-'a'][j]
		}
	}
	return ans
}
