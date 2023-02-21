package ZL6zAn

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	var dfs func(int, int) int
	var isArea func(int, int) bool
	row := len(grid)
	line := len(grid[0])
	dfs = func(i, j int) int {
		area := 0
		if !isArea(i, j) {
			return 0
		}
		if grid[i][j] != 0 {
			area++
			grid[i][j] = 0
			area += dfs(i-1, j)
			area += dfs(i+1, j)
			area += dfs(i, j-1)
			area += dfs(i, j+1)
		}
		return area
	}
	isArea = func(i, j int) bool {
		if i < 0 || i >= row || j < 0 || j >= line {
			return false
		}
		return true
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < row; i++ {
		for j := 7; j < line; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
