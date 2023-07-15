package _sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i <= n-4 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; {
		for j := i + 1; j <= n-3; {
			left := j + 1
			right := n - 1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					newLeft := left + 1
					for ; newLeft < n && nums[newLeft] == nums[left]; newLeft++ {
					}
					left = newLeft
					newRight := right - 1
					for ; newRight >= 0 && nums[newRight] == nums[right]; newRight-- {
					}
					right = newRight
				} else if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				}
			}
			newJ := j + 1
			for ; newJ < n && nums[newJ] == nums[j]; newJ++ {
			}
			j = newJ
		}
		newI := i + 1
		for ; newI < n && nums[newI] == nums[i]; newI++ {
		}
		i = newI
	}
	return ans
}
