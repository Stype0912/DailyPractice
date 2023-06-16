package number_of_unequal_triplets_in_array

func unequalTriplets(nums []int) int {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	left := 0
	ans := 0
	for _, value := range cnt {
		right := len(nums) - left - value
		ans += left * value * right
		left += value
	}
	return ans
}
