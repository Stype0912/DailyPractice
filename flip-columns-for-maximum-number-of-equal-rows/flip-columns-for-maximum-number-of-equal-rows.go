package flip_columns_for_maximum_number_of_equal_rows

func maxEqualRowsAfterFlips(matrix [][]int) int {
	row := len(matrix)
	line := len(matrix[0])
	lineMap := map[string]int{}
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < line; j++ {
				matrix[i][j] = matrix[i][j] ^ 1
			}
		}
		tmpStr := ""
		for j := 0; j < line; j++ {
			if matrix[i][j] == 0 {
				tmpStr += "0"
			} else {
				tmpStr += "1"
			}
		}
		lineMap[tmpStr]++
	}
	ans := 0
	for _, value := range lineMap {
		if value > ans {
			ans = value
		}
	}
	return ans
}
