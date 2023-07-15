package _sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i <= len(nums)-3; {
		j := i + 1
		k := len(nums) - 1
		flag := false
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				m := j + 1
				for m < len(nums) && nums[m] == nums[j] {
					m++
				}
				j = m
				m = k - 1
				for m >= 0 && nums[m] == nums[k] {
					m--
				}
				k = m
				flag = true

			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--

			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			}
		}
		if flag {
			m := i + 1
			for m < len(nums) && nums[m] == nums[i] {
				m++
			}
			i = m
		} else {
			i++
		}
	}
	return ans
}
