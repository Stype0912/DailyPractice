package make_array_zero_by_subtracting_equal_amounts

func minimumOperations(nums []int) int {
	numsMap := map[int]int{}
	for _, num := range nums {
		if num != 0 {
			numsMap[num]++
		}
	}
	return len(numsMap)
}
