package special_array_i

func isArraySpecial(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if (nums[i-1]^nums[i])&1 == 1 {
			continue
		} else {
			return false
		}
	}
	return true
}
