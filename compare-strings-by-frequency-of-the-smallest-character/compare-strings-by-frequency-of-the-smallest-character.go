package compare_strings_by_frequency_of_the_smallest_character

func numSmallerByFrequency(queries []string, words []string) []int {
	ansOfWords := []int{}
	ans := make([]int, len(queries))
	for _, word := range words {
		ansOfWords = append(ansOfWords, f(word))
	}
	for i, query := range queries {
		freq := f(query)
		for _, ansOfWord := range ansOfWords {
			if freq < ansOfWord {
				ans[i]++
			}
		}
	}
	return ans
}

func f(s string) int {
	var minLetter int32 = 'z' + 1
	counterMap := map[int32]int{}
	for _, letter := range s {
		counterMap[letter]++
		if letter < minLetter {
			minLetter = letter
		}
	}
	return counterMap[minLetter]
}
