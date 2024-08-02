package right_triangles

func numberOfRightTriangles(grid [][]int) int64 {
	row := len(grid)
	line := len(grid[0])
	rowSum := make([]int, row)
	lineSum := make([]int, line)
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			rowSum[i] += grid[i][j]
			lineSum[j] += grid[i][j]
		}
	}
	ans := int64(0)
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == 1 && rowSum[i] > 1 && lineSum[j] > 1 {
				ans += int64((rowSum[i] - 1) * (lineSum[j] - 1))
			}
		}
	}
	return ans
}
