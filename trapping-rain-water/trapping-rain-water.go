package trapping_rain_water

func trap(height []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}
