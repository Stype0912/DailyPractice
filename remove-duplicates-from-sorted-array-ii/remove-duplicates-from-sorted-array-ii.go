package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	n := len(nums)
	slow, fast := 2, 2
	if n <= 2 {
		return n
	}
	for ; fast < n; fast++ {
		if nums[fast] != nums[fast-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
