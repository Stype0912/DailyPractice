package moving_stones_until_consecutive

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func numMovesStones(a int, b int, c int) []int {
	tmpa := min(a, min(b, c))
	tmpc := max(a, max(b, c))
	tmpb := a + b + c - tmpa - tmpc
	ans := make([]int, 2)
	if tmpb-tmpa == 1 && tmpc-tmpb == 1 {
		ans[0] = 0
	} else if tmpb-tmpa == 2 || tmpc-tmpb == 2 {
		ans[0] = 1
	} else {
		ans[0] = 2
	}
	ans[1] = tmpc - tmpa - 2
	return ans
}
