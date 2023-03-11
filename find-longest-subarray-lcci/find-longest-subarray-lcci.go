package find_longest_subarray_lcci

func findLongestSubarray(array []string) []string {
	valueIndexMap := map[int]int{}
	valueFirstIndexMap := map[int]int{}
	counter := 0
	for index, arrayEle := range array {
		if (arrayEle >= "A" && arrayEle <= "Z") || (arrayEle >= "a" && arrayEle <= "z") {
			counter++
		} else {
			counter--
		}
		if _, ok := valueFirstIndexMap[counter]; !ok {
			valueFirstIndexMap[counter] = index
		}
		valueIndexMap[counter] = index
	}
	ansIndex := [2]int{0, 0}
	for value, index := range valueIndexMap {
		if value == 0 && index+1 >= ansIndex[1]-ansIndex[0] {
			//fmt.Println(ansIndex[0], ansIndex[1])
			ansIndex = [2]int{-1, index}
		} else {
			if firstIndex, ok := valueFirstIndexMap[value]; (ok && index-firstIndex > ansIndex[1]-ansIndex[0]) ||
				(index-firstIndex == ansIndex[1]-ansIndex[0] && firstIndex < ansIndex[0]) {
				//fmt.Println(firstIndex, index, ansIndex[0])
				ansIndex = [2]int{firstIndex, index}
			}
		}
	}
	ans := []string{}
	for i := ansIndex[0] + 1; i <= ansIndex[1]; i++ {
		ans = append(ans, array[i])
	}
	return ans
}
