package maximum_number_of_pairs_in_array

func numberOfPairs(nums []int) []int {
	numCount := map[int]int{}
	res1, res2 := 0, 0
	for _, num := range nums {
		numCount[num]++
	}
	for num := range numCount {
		res1 += numCount[num] / 2
		if numCount[num]%2 != 0 {
			res2++
		}
	}
	return []int{res1, res2}
}
