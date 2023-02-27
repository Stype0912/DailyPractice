package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		j := i + 2
		for {
			if j >= n || nums[j] != nums[i] {
				break
			}
			if nums[j] == nums[i] {
				nums = append(nums[:j], nums[j+1:]...)
				n--
			}
		}
	}
	return len(nums)
}
