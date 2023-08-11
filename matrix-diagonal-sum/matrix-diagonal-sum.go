package matrix_diagonal_sum

func diagonalSum(mat [][]int) int {
	rows := len(mat)
	sum := 0
	for i := 0; i < rows; i++ {
		sum += mat[i][i]
	}
	for i, j := 0, rows-1; i < rows; i, j = i+1, j-1 {
		if i != j {
			sum += mat[i][j]
		}
	}
	return sum
}
