package all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	var dfs func(int, []int)
	dfs = func(now int, path []int) {
		pathTmp := make([]int, len(path))
		copy(pathTmp, path)
		pathTmp = append(pathTmp, now)
		if now == len(graph)-1 {
			ans = append(ans, pathTmp)
			return
		}
		for i := 0; i < len(graph[now]); i++ {
			//pathTmp := make([]int, len(path))
			//copy(pathTmp, path)
			dfs(graph[now][i], pathTmp)
		}
	}
	dfs(0, []int{})
	return ans
}
