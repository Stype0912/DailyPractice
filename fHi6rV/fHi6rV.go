package fHi6rV

func flipChess(chessboard []string) int {
	row, line := len(chessboard), len(chessboard[0])
	var bfs func(int, int) int
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	bfs = func(i int, j int) int {
		res := 0
		queue := [][]int{{i, j}}
		g := make([][]byte, row)
		for i := range g {
			g[i] = make([]byte, line)
			copy(g[i], chessboard[i])
		}
		g[i][j] = 'X'

		for len(queue) != 0 {
			i, j = queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, direction := range directions {
				x, y := i+direction[0], j+direction[1]
				for x >= 0 && x < row && y >= 0 && y < line && g[x][y] == 'O' {
					x += direction[0]
					y += direction[1]
				}
				if x >= 0 && x < row && y >= 0 && y < line && g[x][y] == 'X' {
					x -= direction[0]
					y -= direction[1]
					res += max(abs(x-i), abs(y-j))
					for x != i || y != j {
						g[x][y] = 'X'
						queue = append(queue, []int{x, y})
						x -= direction[0]
						y -= direction[1]
					}
				}
			}
		}
		return res
	}
	ans := 0
	for i := 0; i < len(chessboard); i++ {
		for j := 0; j < len(chessboard[0]); j++ {
			if chessboard[i][j] == '.' {
				ans = max(ans, bfs(i, j))
			}
		}
	}
	return ans
}
