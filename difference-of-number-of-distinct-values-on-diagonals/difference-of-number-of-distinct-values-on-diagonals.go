package difference_of_number_of_distinct_values_on_diagonals

func differenceOfDistinctValues(grid [][]int) [][]int {
	row := len(grid)
	line := len(grid[0])
	ans := make([][]int, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]int, line)
	}
	for i := 0; i < row; i++ {
		k := i
		j := 0
		leftCounterMap := map[int]int{}
		rightCounterMap := map[int]int{}
		for k < row && j < line {
			rightCounterMap[grid[k][j]]++
			k++
			j++
		}
		k = i
		j = 0
		for k < row && j < line {
			rightCounterMap[grid[k][j]]--
			if rightCounterMap[grid[k][j]] == 0 {
				delete(rightCounterMap, grid[k][j])
			}
			ans[k][j] = abs(len(leftCounterMap), len(rightCounterMap))
			leftCounterMap[grid[k][j]]++
			k++
			j++
		}
	}
	for j := 1; j < line; j++ {
		k := j
		i := 0
		leftCounterMap := map[int]int{}
		rightCounterMap := map[int]int{}
		for i < row && k < line {
			rightCounterMap[grid[i][k]]++
			k++
			i++
		}
		k = j
		i = 0
		for i < row && k < line {
			rightCounterMap[grid[i][k]]--
			if rightCounterMap[grid[i][k]] == 0 {
				delete(rightCounterMap, grid[i][k])
			}
			ans[i][k] = abs(len(leftCounterMap), len(rightCounterMap))
			leftCounterMap[grid[i][k]]++
			k++
			i++
		}
	}
	return ans
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
