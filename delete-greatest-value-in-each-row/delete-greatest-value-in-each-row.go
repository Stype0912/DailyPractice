package delete_greatest_value_in_each_row

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ans := 0
	for i := 0; i < len(grid[0]); i++ {
		tmp := 0
		for j := 0; j < len(grid); j++ {
			tmp = max(tmp, grid[j][i])
		}
		ans += tmp
	}
	return ans
}
