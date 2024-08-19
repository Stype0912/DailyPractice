package find_point_in_most_intervals

import "sort"

func findPointInMostIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 0
	tmp := 0
	maxRight := -1
	for _, p := range intervals {
		if p[0] > maxRight {
			tmp = 1
			maxRight = p[1]
		} else {
			tmp++
		}
		ans = max(ans, tmp)
	}
	return ans
}
