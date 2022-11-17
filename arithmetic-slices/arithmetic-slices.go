package main

func numberOfArithmeticSlices(nums []int) int {
	ansInt := 0
	if len(nums) <= 2 {
		return 0
	}
	ans := make([]int, len(nums)-2)
	if nums[1]*2 == nums[0]+nums[2] {
		ans[0] = 1
	}
	for i := 3; i < len(nums); i++ {
		if nums[i-1]*2 == nums[i-2]+nums[i] {
			ans[i-2] += ans[i-3] + 1
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		ansInt += ans[i]
	}
	return ansInt
}