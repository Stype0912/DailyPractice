package number_of_islands

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if !inArea(n, m, i, j) {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i-1, j)
			dfs(i, j-1)
			dfs(i, j+1)
			dfs(i+1, j)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

func inArea(n, m, r, c int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}
