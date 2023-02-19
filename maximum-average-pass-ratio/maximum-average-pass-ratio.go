package maximum_average_pass_ratio

import (
	"container/heap"
)

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return (h[j][1]+1)*h[j][1]*(h[i][1]-h[i][0]) > (h[i][1]+1)*h[i][1]*(h[j][1]-h[j][0])
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Pop() (_ interface{}) { return }

func (h IntHeap) Push(x interface{}) {}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := IntHeap{}
	for _, class := range classes {
		h = append(h, class)
	}
	heap.Init(&h)
	for i := 0; i < extraStudents; i++ {
		h[0][0]++
		h[0][1]++
		heap.Fix(&h, 0)
	}
	ans := 0.0
	for _, item := range h {
		ans += float64(item[0]) / float64(item[1])
	}
	return ans / float64(len(classes))
}
