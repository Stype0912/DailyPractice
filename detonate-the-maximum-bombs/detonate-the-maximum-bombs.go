package detonate_the_maximum_bombs

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	visited := make([]bool, n)
	ans := 1
	var dfs func(id int)
	var cnt int = 0
	distance := func(i, j, x, y int) int {
		return (i-x)*(i-x) + (j-y)*(j-y)
	}
	dfs = func(id int) {
		if visited[id] {
			return
		}
		visited[id] = true
		cnt++
		ans = max(ans, cnt)
		i, j, r := bombs[id][0], bombs[id][1], bombs[id][2]
		for k := 0; k < n; k++ {
			if !visited[k] && distance(i, j, bombs[k][0], bombs[k][1]) <= r*r {
				dfs(k)
			}
		}
	}
	for i := 0; i < n; i++ {
		visited = make([]bool, n)
		cnt = 0
		dfs(i)
	}
	return ans
}
