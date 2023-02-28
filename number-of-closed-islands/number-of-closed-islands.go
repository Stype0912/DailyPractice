package number_of_closed_islands

func closedIsland(grid [][]int) int {
	row := len(grid)
	line := len(grid[0])
	isArea := func(i, j int) bool {
		if i < 0 || i > row-1 || j < 0 || j > line-1 {
			return false
		} else {
			return true
		}
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if !isArea(i, j) {
			return false
		}
		if grid[i][j] == 1 {
			return true
		} else {
			grid[i][j] = 1
			bool1 := dfs(i-1, j)
			bool2 := dfs(i+1, j)
			bool3 := dfs(i, j-1)
			bool4 := dfs(i, j+1)
			return bool1 && bool2 && bool3 && bool4
		}
	}
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == 0 && dfs(i, j) {
				ans++
			}
		}
	}
	return ans
}
