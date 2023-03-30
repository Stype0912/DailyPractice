package aseY1I

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxProduct(words []string) int {
	wordsMap := []int32{}
	for _, word := range words {
		tmp := int32(0)
		for _, letter := range word {
			tmp |= 1 << (letter - 'a')
		}
		wordsMap = append(wordsMap, tmp)
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if wordsMap[i]&wordsMap[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}
