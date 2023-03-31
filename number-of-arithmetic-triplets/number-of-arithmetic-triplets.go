package number_of_arithmetic_triplets

func arithmeticTriplets(nums []int, diff int) int {
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	ans := 0
	for _, num := range nums {
		if numsMap[num+diff] && numsMap[num+2*diff] {
			ans++
		}
	}
	return ans
}
