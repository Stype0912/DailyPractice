package number_of_provinces

func findCircleNum(isConnected [][]int) int {
	capital := make([]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		capital[i] = i
	}
	var union func(int, int)
	var find func(int) int
	find = func(x int) int {
		if capital[x] == x {
			return x
		}
		return find(capital[x])
	}
	union = func(x, y int) {
		if find(x) != find(y) {
			capital[find(y)] = find(x)
		}
	}
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	capitalMap := map[int]int{}
	for i := 0; i < len(capital); i++ {
		capitalMap[find(i)]++
	}
	return len(capitalMap)
}
