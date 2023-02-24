package shortest_bridge

func shortestBridge(grid [][]int) int {
	var dfs func(int, int)
	n := len(grid)
	isArea := func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= n {
			return false
		}
		return true
	}
	island := [][]int{}
	dfs = func(i int, j int) {
		if !isArea(i, j) {
			return
		}
		if grid[i][j] != 1 {
			return
		}
		grid[i][j] = 2
		island = append(island, []int{i, j})
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		return
	}
	searchIsland := func() {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					dfs(i, j)
					return
				}
			}
		}
	}
	searchIsland()
	param := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i <= 2*n-3; i++ {
		newIsland := [][]int{}
		for _, islandItem := range island {
			row, line := islandItem[0], islandItem[1]
			for _, paramItem := range param {
				newRow := row + paramItem[0]
				newLine := line + paramItem[1]
				if isArea(newRow, newLine) {
					switch grid[newRow][newLine] {
					case 2:
						continue
					case 1:
						return i
					case 0:
						newIsland = append(newIsland, []int{newRow, newLine})
						grid[newRow][newLine] = 2
					}
				}
			}
		}
		island = newIsland
	}
	return 0
}
