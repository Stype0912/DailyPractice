package combination_sum_ii

import "sort"

type intList []int

func (m intList) Len() int {
	return len(m)
}

func (m intList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func combinationSum2(candidates []int, target int) [][]int {
	var dfs func(int, []int, []int, int)
	ans := [][]int{}
	sort.Sort(intList(candidates))
	dfs = func(subTarget int, path []int, subCandidates []int, startIndex int) {
		if subTarget < 0 {
			return
		}
		if subTarget == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i, _ := range subCandidates {
			copySubCandidates := make([]int, len(subCandidates))
			copy(copySubCandidates, subCandidates)
			subCandidate := copySubCandidates[i]
			if subTarget-subCandidate < 0 {
				break
			}
			if i > startIndex && copySubCandidates[i] == copySubCandidates[i-1] {
				continue
			} else if i >= startIndex {
				dfs(subTarget-subCandidate, append(path, subCandidate), append(copySubCandidates[:i], copySubCandidates[i+1:]...), i)
			}
		}
	}
	dfs(target, []int{}, candidates, 0)
	return ans
}
