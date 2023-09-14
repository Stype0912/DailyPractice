package queens_that_can_attack_the_kin

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	mapQueens := map[[2]int]bool{}
	for _, queen := range queens {
		mapQueens[[2]int{queen[0], queen[1]}] = true
	}
	canAttackQueens := map[[2]int]bool{}

	isArea := func(i, j int) bool {
		if i < 0 || i > 7 || j < 0 || j > 7 {
			return false
		}
		return true
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var dfs func(direction int, i, j int)
	dfs = func(direction int, i, j int) {
		nowI := i + directions[direction][0]
		nowJ := j + directions[direction][1]
		if !isArea(i, j) {
			return
		}
		if mapQueens[[2]int{nowI, nowJ}] {
			canAttackQueens[[2]int{nowI, nowJ}] = true
			return
		} else {
			dfs(direction, nowI, nowJ)
		}
		return
	}
	for i := 0; i < 8; i++ {
		dfs(i, king[0], king[1])
	}
	ans := [][]int{}
	for _, queen := range queens {
		if canAttackQueens[[2]int{queen[0], queen[1]}] {
			ans = append(ans, queen)
		}
	}
	return ans
}
