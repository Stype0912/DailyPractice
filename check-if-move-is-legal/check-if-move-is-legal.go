package check_if_move_is_legal

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	var dfs func(i, j int, direction, length int) bool
	n, m := len(board), len(board[0])
	inArea := func(i, j int) bool {
		if i >= 0 && i < n && j >= 0 && j < m {
			return true
		}
		return false
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	dfs = func(i, j int, direction, length int) bool {
		if !inArea(i, j) {
			return false
		}
		if board[i][j] == color && length >= 3 {
			return true
		}
		if board[i][j] == color && length == 2 {
			return false
		}
		if (i == rMove && j == cMove) || (color == 'W' && board[i][j] == 'B') || (color == 'B' && board[i][j] == 'W') {
			return dfs(i+directions[direction][0], j+directions[direction][1], direction, length+1)
		}
		return false
	}
	for i := 0; i < len(directions); i++ {
		if dfs(rMove, cMove, i, 1) {
			return true
		}
	}
	return false
}
