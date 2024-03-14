package largest_element_in_an_array_after_merge_operations

func maxArrayValue(nums []int) int64 {
	sum := int64(nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		if int64(nums[i]) <= sum {
			sum = int64(nums[i]) + sum
		} else {
			sum = int64(nums[i])
		}
	}
	return sum
}
