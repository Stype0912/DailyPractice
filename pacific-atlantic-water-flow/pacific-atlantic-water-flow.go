package pacific_atlantic_water_flow

func pacificAtlantic(heights [][]int) [][]int {
	row := len(heights)
	line := len(heights[0])
	ans := [][]int{}
	pacific := map[int]map[int]bool{}
	atlantic := map[int]map[int]bool{}
	for i := 0; i < row; i++ {
		pacific[i] = map[int]bool{}
		atlantic[i] = map[int]bool{}
	}
	isArea := func(i, j int) bool {
		if i < 0 || i >= row || j < 0 || j >= line {
			return false
		}
		return true
	}
	var dfs func(int, int, map[int]map[int]bool)
	dfs = func(i int, j int, ocean map[int]map[int]bool) {
		if !isArea(i, j) {
			return
		}
		if ocean[i][j] {
			return
		}
		ocean[i][j] = true
		tmp := heights[i][j]
		if isArea(i, j+1) && heights[i][j+1] >= tmp {
			dfs(i, j+1, ocean)
		}
		if isArea(i-1, j) && heights[i-1][j] >= tmp {
			dfs(i-1, j, ocean)
		}
		if isArea(i+1, j) && heights[i+1][j] >= tmp {
			dfs(i+1, j, ocean)
		}
		if isArea(i, j-1) && heights[i][j-1] >= tmp {
			dfs(i, j-1, ocean)
		}
	}
	for i := 0; i < row; i++ {
		dfs(i, 0, pacific)
		dfs(i, line-1, atlantic)
	}
	for j := 0; j < line; j++ {
		dfs(0, j, pacific)
		dfs(row-1, j, atlantic)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			_, ok1 := pacific[i][j]
			_, ok2 := atlantic[i][j]
			if ok1 && ok2 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
