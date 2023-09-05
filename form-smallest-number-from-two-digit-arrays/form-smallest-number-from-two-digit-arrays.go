package form_smallest_number_from_two_digit_arrays

func minNumber(nums1 []int, nums2 []int) int {
	nums1Map := map[int]int{}
	min1 := 10
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for _, num := range nums1 {
		nums1Map[num]++
		min1 = min(min1, num)
	}
	min2 := 10
	ans := 10
	for _, num := range nums2 {
		if _, ok := nums1Map[num]; ok {
			ans = min(ans, num)
		}
		min2 = min(min2, num)
	}
	if ans != 10 {
		return ans
	}
	second := min(min1, min2)
	return 10*second + min1 + min2 - second
}
