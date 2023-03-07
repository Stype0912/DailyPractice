package count_sub_islands

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	row := len(grid2)
	line := len(grid2[0])
	var dfs func(int, int) bool
	dfs = func(i int, j int) bool {
		grid2[i][j] = 0
		var bool1, bool2, bool3, bool4 = true, true, true, true
		if i > 0 && grid2[i-1][j] == 1 {
			bool1 = dfs(i-1, j)
		}
		if i < row-1 && grid2[i+1][j] == 1 {
			bool2 = dfs(i+1, j)
		}
		if j > 0 && grid2[i][j-1] == 1 {
			bool3 = dfs(i, j-1)
		}
		if j < line-1 && grid2[i][j+1] == 1 {
			bool4 = dfs(i, j+1)
		}
		return bool1 && bool2 && bool3 && bool4 && grid1[i][j] != 0
	}
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid2[i][j] == 1 {
				if dfs(i, j) {
					ans++
				}
			}
		}
	}
	return ans
}
