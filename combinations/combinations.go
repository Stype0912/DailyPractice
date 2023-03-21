package combinations

func combine(n int, k int) [][]int {
	ans := [][]int{}
	candidates := []int{}
	for i := 1; i <= n; i++ {
		candidates = append(candidates, i)
	}
	var dfs func(int, []int)
	dfs = func(index int, paths []int) {
		if len(paths) == k {
			ans = append(ans, paths)
			return
		}
		for i := index; i < n; i++ {
			tmp := make([]int, len(paths))
			copy(tmp, paths)
			dfs(i+1, append(tmp, candidates[i]))
		}
	}
	dfs(0, []int{})
	return ans
}
