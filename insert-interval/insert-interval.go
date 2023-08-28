package insert_interval

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			left := ans[len(ans)-1][0]
			right := max(intervals[i][1], ans[len(ans)-1][1])
			ans = ans[:len(ans)-1]
			ans = append(ans, []int{left, right})
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
