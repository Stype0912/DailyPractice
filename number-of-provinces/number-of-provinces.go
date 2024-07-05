package number_of_provinces

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		return find(parent[i])
	}
	union := func(i, j int) {
		if find(i) != find(j) {
			parent[find(i)] = parent[find(j)]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	unionMap := map[int]bool{}
	for _, num := range parent {
		unionMap[find(num)] = true
	}
	return len(unionMap)
}
