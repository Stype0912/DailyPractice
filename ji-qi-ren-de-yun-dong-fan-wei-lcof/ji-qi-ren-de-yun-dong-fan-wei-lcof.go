package ji_qi_ren_de_yun_dong_fan_wei_lcof

func movingCount(m int, n int, k int) int {
	var dfs func(int, int)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	ans := 0
	dfs = func(i, j int) {
		if !isValid(i, j, k) {
			return
		}
		visited[i][j] = true
		ans += 1
		if i > 0 && !visited[i-1][j] {
			dfs(i-1, j)
		}
		if i < m-1 && !visited[i+1][j] {
			dfs(i+1, j)
		}
		if j > 0 && !visited[i][j-1] {
			dfs(i, j-1)
		}
		if j < n-1 && !visited[i][j+1] {
			dfs(i, j+1)
		}
		visited[i][j] = false
	}
	dfs(0, 0)
	return ans
}

func isValid(m, n, k int) bool {
	tmp := 0
	if m == 100 {
		tmp += 1
	} else {
		tmp += m / 10
		tmp += m % 10
	}
	if n == 100 {
		tmp += 1
	} else {
		tmp += n / 10
		tmp += n % 10
	}
	return tmp <= k
}
