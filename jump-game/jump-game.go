package jump_game

func canJump(nums []int) bool {
	n := len(nums)
	canReach := 0
	for i := 0; i < n; i++ {
		if canReach < i {
			return false
		}
		canReach = max(canReach, i+nums[i])
		if canReach >= n-1 {
			return true
		}
	}
	return false
}
