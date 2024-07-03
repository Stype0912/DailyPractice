package rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	numsNew := append(nums[n-k:], nums[:n-k]...)
	copy(nums, numsNew)
}
