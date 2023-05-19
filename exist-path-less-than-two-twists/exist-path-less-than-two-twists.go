package exist_path_less_than_two_twists

func pathExists(matrix [][]int) bool {
	row := len(matrix)
	line := len(matrix[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, line)
	}
	var dfs func(i, j int, direction int, twistCnt int) bool
	dfs = func(i, j int, direction int, twistCnt int) bool {
		if visited[i][j] {
			return false
		}
		visited[i][j] = true
		if matrix[i][j] == 1 && twistCnt <= 2 {
			return true
		}
		if twistCnt > 2 || matrix[i][j] == 2 {
			return false
		}
		var bool1, bool2, bool3, bool4 bool
		switch direction {
		case 0:
			if i > 0 {
				bool1 = dfs(i-1, j, 1, 0)
			}
			if i < row-1 {
				bool2 = dfs(i+1, j, 2, 0)
			}
			if j > 0 {
				bool3 = dfs(i, j-1, 3, 0)
			}
			if j < line-1 {
				bool4 = dfs(i, j+1, 4, 0)
			}
		case 1:
			if i > 0 {
				bool1 = dfs(i-1, j, 1, twistCnt)
			}
			if i < row-1 {
				bool2 = dfs(i+1, j, 2, twistCnt+1)
			}
			if j < line-1 {
				bool4 = dfs(i, j+1, 4, twistCnt+1)
			}
			if j > 0 {
				bool3 = dfs(i, j-1, 3, twistCnt+1)
			}
		case 2:
			if i > 0 {
				bool1 = dfs(i-1, j, 1, twistCnt+1)
			}
			if i < row-1 {
				bool2 = dfs(i+1, j, 2, twistCnt)
			}
			if j > 0 {
				bool3 = dfs(i, j-1, 3, twistCnt+1)
			}
			if j < line-1 {
				bool4 = dfs(i, j+1, 4, twistCnt+1)
			}
		case 3:
			if i > 0 {
				bool1 = dfs(i-1, j, 1, twistCnt+1)
			}
			if i < row-1 {
				bool2 = dfs(i+1, j, 2, twistCnt+1)
			}
			if j > 0 {
				bool3 = dfs(i, j-1, 3, twistCnt)
			}
			if j < line-1 {
				bool4 = dfs(i, j+1, 4, twistCnt+1)
			}
		case 4:
			if i > 0 {
				bool1 = dfs(i-1, j, 1, twistCnt+1)
			}
			if i < row-1 {
				bool2 = dfs(i+1, j, 2, twistCnt+1)
			}
			if j > 0 {
				bool3 = dfs(i, j-1, 3, twistCnt+1)
			}
			if j < line-1 {
				bool4 = dfs(i, j+1, 4, twistCnt)
			}
		}
		visited[i][j] = false
		return bool1 || bool2 || bool3 || bool4
	}
	return dfs(1, 1, 0, 0)
}
