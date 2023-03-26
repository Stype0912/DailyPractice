package find_subarrays_with_equal_sum

func findSubarrays(nums []int) bool {
	sum := map[int]bool{}
	for i := 0; i < len(nums)-1; i++ {
		if _, ok := sum[nums[i]+nums[i+1]]; ok {
			return true
		} else {
			sum[nums[i]+nums[i+1]] = true
		}
	}
	return false
}
