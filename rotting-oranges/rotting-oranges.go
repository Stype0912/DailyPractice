package rotting_oranges

func orangesRotting(grid [][]int) (res int) {
	row := len(grid)
	line := len(grid[0])
	count := 0
	var queue [][]int
	res = 0
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == 1 {
				count++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	for {
		if len(queue) == 0 || count == 0 {
			break
		}
		n := len(queue)
		for i := 0; i < n; i++ {
			element := queue[0]
			queue = queue[1:]
			if element[0]-1 >= 0 && grid[element[0]-1][element[1]] == 1 {
				grid[element[0]-1][element[1]] = 2
				count--
				queue = append(queue, []int{element[0] - 1, element[1]})
			}
			if element[0]+1 < row && grid[element[0]+1][element[1]] == 1 {
				grid[element[0]+1][element[1]] = 2
				count--
				queue = append(queue, []int{element[0] + 1, element[1]})
			}
			if element[1]-1 >= 0 && grid[element[0]][element[1]-1] == 1 {
				grid[element[0]][element[1]-1] = 2
				count--
				queue = append(queue, []int{element[0], element[1] - 1})
			}
			if element[1]+1 < line && grid[element[0]][element[1]+1] == 1 {
				grid[element[0]][element[1]+1] = 2
				count--
				queue = append(queue, []int{element[0], element[1] + 1})
			}
		}
		res++
	}
	if count > 0 {
		return -1
	} else {
		return res
	}
}
