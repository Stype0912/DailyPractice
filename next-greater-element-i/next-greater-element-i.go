package next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	var stack, res []int
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	for _, item := range nums1 {
		res = append(res, mp[item])
	}
	return res
}
