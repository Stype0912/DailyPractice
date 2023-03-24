package minimum_average_difference

import "math"

func minimumAverageDifference(nums []int) int {
	prefix := 0
	suffix := 0
	index := 0
	minAverage := -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		suffix = sum - prefix
		prefixAverage := prefix / (i + 1)
		var suffixAverage int
		if suffix == 0 {
			suffixAverage = 0
		} else {
			suffixAverage = suffix / (len(nums) - i - 1)
		}
		averageDifference := int(math.Abs(float64(prefixAverage - suffixAverage)))
		if minAverage == -1 || averageDifference < minAverage {
			minAverage = averageDifference
			index = i
		}
	}
	return index
}
