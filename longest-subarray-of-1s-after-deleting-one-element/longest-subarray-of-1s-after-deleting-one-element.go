package longest_subarray_of_1s_after_deleting_one_element

func longestSubarray(nums []int) int {
	zeroPosMap := map[int][]int{}
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
			continue
		}
		zeroPosMap[i] = append(zeroPosMap[i], counter)
		counter = 0
	}
	counter = 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 1 {
			counter++
			continue
		}
		zeroPosMap[i] = append(zeroPosMap[i], counter)
		counter = 0
	}
	if len(zeroPosMap) == 0 {
		return len(nums) - 1
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ans := 0
	for _, value := range zeroPosMap {
		ans = max(ans, value[0]+value[1])
	}
	return ans
}
