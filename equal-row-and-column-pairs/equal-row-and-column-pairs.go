package equal_row_and_column_pairs

import "fmt"

func equalPairs(grid [][]int) int {
	lineMap := map[string]int{}
	ans := 0
	for _, line := range grid {
		lineMap[fmt.Sprint(line)]++
	}
	for j := 0; j < len(grid); j++ {
		arrTmp := []int{}
		for i := 0; i < len(grid); i++ {
			arrTmp = append(arrTmp, grid[i][j])
		}
		if count, ok := lineMap[fmt.Sprint(arrTmp)]; ok {
			ans += count
		}
	}
	return ans
}
