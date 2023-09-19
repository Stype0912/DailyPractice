package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row := len(matrix)
	line := len(matrix[0])
	var dfs func(i, j, direction int)
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, line)
	}

	isArea := func(i, j int) bool {
		if i < 0 || i >= row || j < 0 || j >= line {
			return false
		}
		return true
	}
	dfs = func(i, j, direction int) {
		ans = append(ans, matrix[i][j])
		visited[i][j] = true
		if !isArea(i+directions[direction][0], j+directions[direction][1]) || visited[i+directions[direction][0]][j+directions[direction][1]] {
			direction = (direction + 1) % 4
			if !isArea(i+directions[direction][0], j+directions[direction][1]) || visited[i+directions[direction][0]][j+directions[direction][1]] {
				return
			}
		}
		dfs(i+directions[direction][0], j+directions[direction][1], direction)
	}
	dfs(0, 0, 0)
	return ans
}
