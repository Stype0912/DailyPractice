package odd_string_difference

import "reflect"

func oddString(words []string) string {
	//n := len(words[0])
	wordDiff := make([]int, len(words[0]))
	wordDiff1 := getDiff(words[0])
	wordDiff2 := getDiff(words[1])
	if reflect.DeepEqual(wordDiff1, wordDiff2) {
		wordDiff = wordDiff1
		for i := 2; i < len(words); i++ {
			wordDiffTmp := getDiff(words[i])
			if !reflect.DeepEqual(wordDiffTmp, wordDiff) {
				return words[i]
			}
		}
	} else {
		wordDiff3 := getDiff(words[2])
		if reflect.DeepEqual(wordDiff1, wordDiff3) {
			return words[1]
		} else {
			return words[0]
		}
	}
	return ""
}

func getDiff(word string) []int {
	n := len(word)
	ans := make([]int, n-1)
	for j := 1; j < n; j++ {
		ans = append(ans, int(word[j]-word[j-1]))
	}
	return ans
}
