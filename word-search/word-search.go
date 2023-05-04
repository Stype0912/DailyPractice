package word_search

func exist(board [][]byte, word string) bool {
	row := len(board)
	line := len(board[0])
	visited := make([][]bool, row)
	for k := 0; k < row; k++ {
		visited[k] = make([]bool, line)
	}
	var dfs func([][]byte, int, int, string) bool
	dfs = func(boardTmp [][]byte, i int, j int, s string) bool {
		// 搜完了返回true
		if len(s) == 0 {
			return true
		}
		visited[i][j] = true
		var bool1, bool2, bool3, bool4 bool
		if i > 0 && boardTmp[i-1][j] == s[0] && !visited[i-1][j] {
			bool1 = dfs(boardTmp, i-1, j, s[1:])
		}
		if i < row-1 && boardTmp[i+1][j] == s[0] && !visited[i+1][j] {
			bool2 = dfs(boardTmp, i+1, j, s[1:])
		}
		if j > 0 && boardTmp[i][j-1] == s[0] && !visited[i][j-1] {
			bool3 = dfs(boardTmp, i, j-1, s[1:])
		}
		if j < line-1 && boardTmp[i][j+1] == s[0] && !visited[i][j+1] {
			bool4 = dfs(boardTmp, i, j+1, s[1:])
		}
		visited[i][j] = false
		return bool1 || bool2 || bool3 || bool4
	}
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if board[i][j] == word[0] {
				visited = make([][]bool, row)
				for k := 0; k < row; k++ {
					visited[k] = make([]bool, line)
				}
				if dfs(board, i, j, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}
