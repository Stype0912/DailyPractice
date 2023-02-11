package number_of_islands

func numIslands(grid [][]byte) int {
	res := 0
	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, r, c int) {
		if !inArea(grid, r, c) {
			return
		}
		if grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
		return
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func inArea(grid [][]byte, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
