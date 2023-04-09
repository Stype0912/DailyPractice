package check_distances_between_same_letters

func checkDistances(s string, distance []int) bool {
	letterMap := map[int32]int{}
	for i, letter := range s {
		if index, ok := letterMap[letter]; !ok {
			letterMap[letter] = i
		} else {
			if i-index-1 != distance[letter-'a'] {
				return false
			}
		}
	}
	return true
}
