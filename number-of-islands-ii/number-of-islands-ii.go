package number_of_islands_ii

func numIslands2(m, n int, positions [][]int) (ans []int) {
	length := m * n
	parent := make([]int, length)
	for i := 0; i < length; i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] == x {
			return x
		}
		return find(parent[x])
	}
	union := func(x, y int) bool {
		if find(x) != find(y) {
			parent[find(x)] = parent[find(y)]
			return true
		}
		return false
	}
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	cnt := 0
	dirs := [5]int{-1, 0, 1, 0, -1}
	for _, p := range positions {
		i, j := p[0], p[1]
		if grid[i][j] == 1 {
			ans = append(ans, cnt)
			continue
		}
		grid[i][j] = 1
		cnt++
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && union(i*n+j, x*n+y) {
				cnt--
			}
		}
		ans = append(ans, cnt)
	}
	return
}
