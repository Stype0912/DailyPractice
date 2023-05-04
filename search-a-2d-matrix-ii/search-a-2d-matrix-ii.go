package search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	line := len(matrix[0])
	i := 0
	j := line - 1
	for {
		if i >= row || j < 0 {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if target > matrix[i][j] {
			i++
			continue
		}
		if target < matrix[i][j] {
			j--
			continue
		}
	}
}
