package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(int, []int, int)
	ans := [][]int{}
	dfs = func(subTarget int, path []int, startIndex int) {
		if subTarget < 0 {
			return
		}
		if subTarget == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			ans = append(ans, newPath)
		}
		for i, _ := range candidates {
			if i >= startIndex {
				dfs(subTarget-candidates[i], append(path, candidates[i]), i)
			}
		}
	}
	dfs(target, []int{}, 0)
	return ans
}
