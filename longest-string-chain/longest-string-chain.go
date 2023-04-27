package longest_string_chain

import "sort"

type Words []string

func (w Words) Len() int {
	return len(w)
}

func (w Words) Less(i, j int) bool {
	return len(w[i]) < len(w[j])
}

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestStrChain(words []string) int {
	wordMap := map[string]int{}
	wordsNew := Words(words)
	sort.Sort(wordsNew)
	for i := 0; i < len(wordsNew); i++ {
		wordMap[wordsNew[i]] = 1
	}
	ans := 1
	for i := 0; i < len(wordsNew); i++ {
		tmp := 1
		for j := 0; j < len(wordsNew[i]); j++ {
			preWord := wordsNew[i][:j] + wordsNew[i][j+1:]
			if wordMap[preWord]+1 > tmp {
				tmp = wordMap[preWord] + 1
				wordMap[wordsNew[i]] = tmp
			}
		}
		ans = max(tmp, ans)
	}
	return ans
}
