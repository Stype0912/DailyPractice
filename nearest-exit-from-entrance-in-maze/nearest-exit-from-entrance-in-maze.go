package nearest_exit_from_entrance_in_maze

func nearestExit(maze [][]byte, entrance []int) int {
	stack := [][3]int{{entrance[0], entrance[1], 0}}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	row, line := len(maze), len(maze[0])
	isArea := func(i, j int) bool {
		if i < 0 || i > row-1 || j < 0 || j > line-1 {
			return false
		}
		return true
	}
	for {
		if len(stack) == 0 {
			break
		}
		now := stack[0]
		stack = stack[1:]
		maze[now[0]][now[1]] = '+'
		if (now[0] == 0 || now[0] == row-1 || now[1] == 0 || now[1] == line-1) && !(now[0] == entrance[0] && now[1] == entrance[1]) {
			return now[2]
		}
		for _, direction := range directions {
			i, j := now[0]+direction[0], now[1]+direction[1]
			if isArea(i, j) && maze[i][j] == '.' {
				stack = append(stack, [3]int{i, j, now[2] + 1})
				maze[i][j] = '+'
				//fmt.Println(i, j, now[2]+1)
			}
		}
	}
	return -1
}
