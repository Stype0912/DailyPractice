package frequency_of_the_most_frequent_element

import "sort"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	left := 0
	right := 1
	total := 0
	for right = 1; right < len(nums); right++ {
		total += (nums[right] - nums[right-1]) * (right - left)
		for total > k {
			total -= nums[right] - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
