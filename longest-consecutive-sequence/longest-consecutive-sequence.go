package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num] = 1
	}
	maxLength := 0
	for key, _ := range numMap {
		if _, ok := numMap[key+maxLength]; !ok {
			continue
		}
		for i := key + 1; ; i++ {
			if _, ok := numMap[i]; !ok {
				maxLength = max(maxLength, i-key)
				break
			}
		}
	}
	return maxLength
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
