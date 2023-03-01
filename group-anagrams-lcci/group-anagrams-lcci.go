package group_anagrams_lcci

func groupAnagrams(strs []string) [][]string {
	wordMap := map[[26]int][]int{}
	for i, str := range strs {
		counter := [26]int{}
		for _, letter := range str {
			counter[letter-'a']++
		}
		wordMap[counter] = append(wordMap[counter], i)
	}
	var ans [][]string
	for _, values := range wordMap {
		var tmp []string
		for _, value := range values {
			tmp = append(tmp, strs[value])
		}
		ans = append(ans, tmp)
	}
	return ans
}
