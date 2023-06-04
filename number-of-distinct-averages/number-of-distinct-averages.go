package number_of_distinct_averages

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	averageMap := map[int]int{}
	for i := 0; i <= len(nums)/2-1; i++ {
		j := len(nums) - 1 - i
		averageMap[nums[i]+nums[j]]++
	}
	return len(averageMap)
}
