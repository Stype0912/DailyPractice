package minimum_time_to_repair_cars

import (
	"container/heap"
	"sort"
)

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][0]*(h[i][1]+1)*(h[i][1]+1) < h[j][0]*(h[j][1]+1)*(h[j][1]+1)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func repairCars(ranks []int, cars int) int64 {
	ans := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	sort.Ints(ranks)
	h := IntHeap{}
	for _, rank := range ranks {
		h = append(h, []int{rank, 0})
	}
	heap.Init(&h)
	for cars > 0 {
		newEle := heap.Pop(&h).([]int)
		rank := newEle[0]
		count := newEle[1]
		ans = max(ans, rank*(count+1)*(count+1))
		heap.Push(&h, []int{rank, count + 1})
		heap.Fix(&h, 0)
		cars--
	}
	return int64(ans)
}
