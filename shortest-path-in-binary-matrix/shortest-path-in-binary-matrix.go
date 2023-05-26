package shortest_path_in_binary_matrix

import (
	"math"
)

//func shortestPathBinaryMatrix(grid [][]int) int {
//	n := len(grid)
//	visit := make([][]bool, n)
//	for i := 0; i < n; i++ {
//		visit[i] = make([]bool, n)
//	}
//	var dfs func([][]int, int, int, int)
//	isArea := func(i, j int) bool {
//		return i >= 0 && i < n && j >= 0 && j < n
//	}
//	min := func(i, j int) int {
//		if i < j {
//			return i
//		}
//		return j
//	}
//	ans := math.MaxInt
//	dfs = func(matrix [][]int, i int, j int, step int) {
//		if step > ans {
//			return
//		}
//		if !isArea(i, j) || visit[i][j] || matrix[i][j] == 1 {
//			return
//		}
//		if i == n-1 && j == n-1 && matrix[i][j] == 0 {
//			ans = min(ans, step)
//		}
//		visit[i][j] = true
//		dfs(matrix, i-1, j-1, step+1)
//		dfs(matrix, i-1, j, step+1)
//		dfs(matrix, i-1, j+1, step+1)
//		dfs(matrix, i, j-1, step+1)
//		dfs(matrix, i, j+1, step+1)
//		dfs(matrix, i+1, j-1, step+1)
//		dfs(matrix, i+1, j, step+1)
//		dfs(matrix, i+1, j+1, step+1)
//		visit[i][j] = false
//	}
//	dfs(grid, 0, 0, 1)
//	if ans == math.MaxInt {
//		return -1
//	} else {
//		return ans
//	}
//}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	visit := make([][]bool, n)
	for i := 0; i < n; i++ {
		visit[i] = make([]bool, n)
	}
	if grid[0][0] == 1 {
		return -1
	}
	isArea := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	nodeList := [][]int{{0, 0, 1}}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ans := math.MaxInt
	for len(nodeList) != 0 {
		//fmt.Println(len(nodeList))
		node := nodeList[0]
		nodeList = nodeList[1:]
		row := node[0]
		line := node[1]
		step := node[2]
		if row == n-1 && line == n-1 {
			ans = min(ans, step)
			return ans
		}
		visit[row][line] = true
		for _, direction := range directions {
			if isArea(row+direction[0], line+direction[1]) && !visit[row+direction[0]][line+direction[1]] && grid[row+direction[0]][line+direction[1]] == 0 {
				visit[row+direction[0]][line+direction[1]] = true
				nodeList = append(nodeList, []int{row + direction[0], line + direction[1], step + 1})
			}
		}
	}
	return -1
}
