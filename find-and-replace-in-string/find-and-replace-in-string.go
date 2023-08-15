package find_and_replace_in_string

import "sort"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n := len(indices)
	isReplacable := make([]bool, n)
	sortedIndices := make([][]int, n)
	for i := 0; i < n; i++ {
		sortedIndices[i] = []int{i, indices[i]}
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return sortedIndices[i][1] < sortedIndices[j][1]
	})
	for i := 0; i < n; i++ {
		for j := 0; j < len(sources[sortedIndices[i][0]]); j++ {
			if sortedIndices[i][1]+j >= len(s) || s[sortedIndices[i][1]+j] != sources[sortedIndices[i][0]][j] {
				break
			}
			if j == len(sources[sortedIndices[i][0]])-1 {
				isReplacable[i] = true
			}
		}
	}
	skew := 0
	for i := 0; i < n; i++ {
		if isReplacable[i] {
			s = s[:sortedIndices[i][1]+skew] + targets[sortedIndices[i][0]] + s[sortedIndices[i][1]+len(sources[sortedIndices[i][0]])+skew:]
			skew += len(targets[sortedIndices[i][0]]) - len(sources[sortedIndices[i][0]])
		}
	}
	return s
}
