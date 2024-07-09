package jump_game_ii

func jump(nums []int) int {
	end := 0
	maxPos := 0
	n := len(nums)
	step := 0
	for i := 0; i < n; i++ {
		maxPos = max(maxPos, i+nums[i])
		if end == i {
			step++
			end = maxPos
		}
	}
	return step
}
