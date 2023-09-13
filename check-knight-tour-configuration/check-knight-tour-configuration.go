package check_knight_tour_configuration

func checkValidGrid(grid [][]int) bool {
	row := len(grid)
	isArea := func(i, j int) bool {
		if i < 0 || i > row-1 || j < 0 || j > row-1 {
			return false
		}
		return true
	}
	directions := [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}
	var dfs func(i, j, val int) bool
	dfs = func(i, j, val int) bool {
		if !isArea(i, j) || grid[i][j] != val+1 {
			return false
		}
		if grid[i][j] == row*row-1 {
			return true
		}
		res := false
		for _, direction := range directions {
			res = res || dfs(i+direction[0], j+direction[1], val+1)
		}
		return res
	}
	return dfs(0, 0, -1)
}
