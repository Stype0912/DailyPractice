package unique_paths_iii

func uniquePathsIII(grid [][]int) int {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	n := 0
	row := len(grid)
	line := len(grid[0])
	starti, startj := 0, 0
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == 0 {
				n++
			}
			if grid[i][j] == 1 {
				n++
				starti = i
				startj = j
			}
		}
	}
	isArea := func(i, j int) bool {
		if i < 0 || i > row-1 || j < 0 || j > line-1 {
			return false
		}
		return true
	}
	var dfs func(i, j, nTmp int) int
	dfs = func(i, j, nTmp int) int {
		if grid[i][j] == 2 {
			if nTmp == 0 {
				return 1
			}
			return 0
		}
		valueTmp := grid[i][j]
		res := 0
		grid[i][j] = -1
		for k := 0; k < 4; k++ {
			ni, nj := i+directions[k][0], j+directions[k][1]
			if isArea(ni, nj) && (grid[ni][nj] == 0 || grid[ni][nj] == 2) {
				res += dfs(ni, nj, nTmp-1)
			}
		}
		grid[i][j] = valueTmp
		return res
	}
	return dfs(starti, startj, n)
}
