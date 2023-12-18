package find_peak_element

import "math"

func findPeakElement(nums []int) int {
	maxNum := math.MinInt
	ans := -1
	for i, num := range nums {
		if num > maxNum {
			maxNum = num
			ans = i
		}
	}
	return ans
}
