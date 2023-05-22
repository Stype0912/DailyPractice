package subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		pathTmp := make([]int, len(path))
		copy(pathTmp, path)
		ans = append(ans, pathTmp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i-1] == nums[i] {
				continue
			}
			dfs(i+1, append(pathTmp, nums[i]))
		}
	}
	dfs(0, []int{})
	return ans
}
