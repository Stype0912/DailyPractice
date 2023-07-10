package _sum_closest

import "math"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	distance := math.MaxInt
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum, target) < distance {
					distance = abs(sum, target)
					ans = sum
				}
			}
		}
	}
	return ans
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
