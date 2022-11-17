package main

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return MaxInt(nums[0], nums[1])
	}
	var ans = []int{nums[0], MaxInt(nums[0], nums[1])}
	for i := 2; i < length; i++ {
		ans = append(ans, MaxInt(nums[i]+ans[i-2], ans[i-1]))
	}
	return ans[length-1]
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}