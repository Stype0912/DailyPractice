package maximum_xor_for_each_query

import (
	"math"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	maxNum := int(math.Pow(float64(2), float64(maximumBit))) - 1
	numsXor := make([]int, n)
	numsXor[n-1] = nums[0]
	for i := 1; i < n; i++ {
		numsXor[n-1-i] = numsXor[n-i] ^ nums[i]
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, maxNum^numsXor[i])
	}
	return ans
}
