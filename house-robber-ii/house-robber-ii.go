package house_robber_ii

func rob(nums []int) int {
	n := len(nums)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp1 := map[int]int{-1: 0, 0: nums[0]}
	dp2 := map[int]int{0: 0, 1: nums[1]}
	for i := 1; i < n-1; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}
	for i := 2; i < n; i++ {
		dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
	}
	return max(dp1[n-2], dp2[n-1])
}
