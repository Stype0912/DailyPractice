package count_servers_that_communicate

func countServers(grid [][]int) int {
	row := len(grid)
	line := len(grid[0])
	label := make([][]int, row)
	for i := 0; i < row; i++ {
		label[i] = make([]int, line)
	}
	serverMap := map[[2]int]int{}
	servers := [][2]int{}
	ans := 0
	for i := 0; i < row; i++ {
		rowCount := 0
		servers = [][2]int{}
		for j := 0; j < line; j++ {
			if grid[i][j] == 1 {
				rowCount++
				servers = append(servers, [2]int{i, j})
			}
		}
		if rowCount >= 2 {
			ans += rowCount
			for _, server := range servers {
				serverMap[server] = 1
			}
		}
	}
	servers = [][2]int{}
	for j := 0; j < line; j++ {
		lineCount := 0
		servers = [][2]int{}
		for i := 0; i < row; i++ {
			if grid[i][j] == 1 {
				lineCount++
				servers = append(servers, [2]int{i, j})
			}
		}
		if lineCount >= 2 {
			for _, server := range servers {
				if serverMap[server] == 1 {
					lineCount--
				}
			}
			ans += lineCount
		}
	}
	return ans
}
