package minimum_interval_to_include_each_query

import (
	"container/heap"
	"sort"
)

type IntervalHeap [][]int

func (h IntervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h IntervalHeap) Len() int {
	return len(h)
}

func (h IntervalHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	sortedQueries := [][]int{}
	for i := 0; i < len(queries); i++ {
		sortedQueries = append(sortedQueries, []int{i, queries[i]})
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][1] < sortedQueries[j][1]
	})
	ans := make([]int, len(sortedQueries))
	intervalHeap := &IntervalHeap{}
	heap.Init(intervalHeap)
	j := 0
	for i := 0; i < len(sortedQueries); i++ {
		value := sortedQueries[i][1]
		for j < len(intervals) && intervals[j][0] <= value {
			heap.Push(intervalHeap, []int{intervals[j][1] - intervals[j][0] + 1, intervals[j][1]})
			j++
		}
		for intervalHeap.Len() > 0 && (*intervalHeap)[0][1] < value {
			heap.Pop(intervalHeap)
		}
		if intervalHeap.Len() > 0 {
			ans[sortedQueries[i][0]] = (*intervalHeap)[0][0]
		} else {
			ans[sortedQueries[i][0]] = -1
		}
	}
	return ans
}
