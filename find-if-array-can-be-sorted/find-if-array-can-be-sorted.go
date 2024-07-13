package find_if_array_can_be_sorted

import "math"

func canSortArray(nums []int) bool {
	n := len(nums)
	currentGroupMax := nums[0]
	currentGroupMin := nums[0]
	lastGroupMax := math.MinInt
	for i := 1; i < n; i++ {
		if countBitOne(nums[i]) == countBitOne(currentGroupMax) {
			currentGroupMax = max(currentGroupMax, nums[i])
			currentGroupMin = min(currentGroupMin, nums[i])
		} else {
			lastGroupMax = currentGroupMax
			currentGroupMax = nums[i]
			currentGroupMin = nums[i]
		}
		if currentGroupMin < lastGroupMax {
			return false
		}
	}
	return true
}

func countBitOne(num int) int {
	ans := 0
	for num != 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}
