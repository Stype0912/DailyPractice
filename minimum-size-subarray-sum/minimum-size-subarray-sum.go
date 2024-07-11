package minimum_size_subarray_sum

import "math"

func minSubArrayLen(target int, nums []int) int {
	if nums[0] >= target {
		return 1
	}
	if len(nums) < 2 {
		return 0
	}
	i, j := 0, 0
	sum := 0
	n := len(nums)
	ans := math.MaxInt
	for j < n {
		sum += nums[j]
		for sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
