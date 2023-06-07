package mice_and_cheese

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	ans := 0
	diffs := make([]int, len(reward1))
	for i := 0; i < len(reward1); i++ {
		ans += reward2[i]
		diffs[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diffs)
	for i := 1; i <= k; i++ {
		ans += diffs[len(reward1)-i]
	}
	return ans
}
