package game_of_life

func gameOfLife(board [][]int) {
	row := len(board)
	line := len(board[0])
	newBoard := make([][]int, row)
	for i := 0; i < row; i++ {
		newBoard[i] = make([]int, line)
	}
	isArea := func(i, j int) bool {
		if i >= 0 && i < row && j >= 0 && j < line {
			return true
		}
		return false
	}
	positions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			live := 0
			for _, pos := range positions {
				x, y := i+pos[0], j+pos[1]
				if !isArea(x, y) {
					continue
				}
				if board[x][y] == 1 {
					live++
				}
			}
			if board[i][j] == 1 {
				if live < 2 || live > 3 {
					newBoard[i][j] = 0
				} else {
					newBoard[i][j] = 1
				}
			} else {
				if live == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			board[i][j] = newBoard[i][j]
		}
	}
}
