package remove_element

func removeElement(nums []int, val int) int {
	ans := 0
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, -1)
			ans++
		} else {
			i++
		}
	}
	return len(nums) - ans
}
