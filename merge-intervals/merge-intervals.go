package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//ans := [][]int{}
	for i := 0; i < len(intervals)-1; {
		//nowStart := intervals[i][0]
		nowEnd := intervals[i][1]
		nextStart := intervals[i+1][0]
		nextEnd := intervals[i+1][1]
		if nextStart > nowEnd {
			i++
		} else {
			newEnd := max(nowEnd, nextEnd)
			intervals[i+1][1] = newEnd
			intervals = append(intervals[:i], intervals[i+1:]...)
		}
	}
	return intervals
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
