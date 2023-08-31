package minimum_degree_of_a_connected_trio_in_a_graph

import "math"

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	degree := map[int]int{}
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0]-1, edges[i][1]-1
		g[x][y] = 1
		g[y][x] = 1
		degree[x]++
		degree[y]++
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if g[i][j] != 1 {
				continue
			}
			for k := j + 1; k < n; k++ {
				if g[j][k] == 1 && g[k][i] == 1 {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
