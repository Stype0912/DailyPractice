package pond_sizes_lcci

import "sort"

func pondSizes(land [][]int) []int {
	var dfs func(int, int, int) int
	row := len(land)
	line := len(land[0])
	isArea := func(i, j int) bool {
		if i < 0 || i > row-1 || j < 0 || j > line-1 {
			return false
		}
		return true
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	dfs = func(i int, j int, square int) int {
		if !isArea(i, j) {
			return square
		}
		if land[i][j] != 0 {
			return square
		}
		square++
		land[i][j] = 1
		for _, direction := range directions {
			square = dfs(i+direction[0], j+direction[1], square)
		}
		return square
	}
	ans := []int{}
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if land[i][j] == 0 {
				ans = append(ans, dfs(i, j, 0))
			}
		}
	}
	sort.Ints(ans)
	return ans
}
