package find_common_elements_between_two_arrays

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]bool{}
	nums2Map := map[int]bool{}
	ans := make([]int, 2)
	for _, num := range nums1 {
		nums1Map[num] = true
	}
	for _, num := range nums2 {
		nums2Map[num] = true
	}
	for i := 0; i < len(nums1); i++ {
		if nums2Map[nums1[i]] {
			ans[0]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if nums1Map[nums2[i]] {
			ans[1]++
		}
	}
	return ans
}
