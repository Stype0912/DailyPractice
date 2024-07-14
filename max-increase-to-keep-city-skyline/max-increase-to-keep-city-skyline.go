package max_increase_to_keep_city_skyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	lineMax := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rowMax[i] = max(rowMax[i], grid[i][j])
			lineMax[j] = max(lineMax[j], grid[i][j])
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += min(rowMax[i], lineMax[j]) - grid[i][j]
		}
	}
	return ans
}
