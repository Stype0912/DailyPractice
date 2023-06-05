package apply_operations_to_an_array

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}
	cnt := 0
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			//nums = append(nums, 0)
			cnt++
		} else {
			i++
		}
	}
	for i := 0; i < cnt; i++ {
		nums = append(nums, 0)
	}
	return nums
}
