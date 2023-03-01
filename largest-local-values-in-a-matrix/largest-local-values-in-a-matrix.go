package largest_local_values_in_a_matrix

func largestLocal(grid [][]int) [][]int {
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
	}
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			tmp := 0
			for k := 0; k < 3; k++ {
				for p := 0; p < 3; p++ {
					tmp = max(tmp, grid[i+k][j+p])
				}
			}
			ans[i][j] = tmp
		}
	}
	return ans
}
