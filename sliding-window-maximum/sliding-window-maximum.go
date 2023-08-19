package sliding_window_maximum

import (
	"container/heap"
)

type Num struct {
	index int
	value int
}
type NumsHeap []Num

func (h NumsHeap) Len() int {
	return len(h)
}

func (h NumsHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}
func (h NumsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NumsHeap) Push(x any) {
	*h = append(*h, x.(Num))
}

func (h *NumsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	h := &NumsHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, Num{
			index: i,
			value: nums[i],
		})
	}
	tmp := heap.Pop(h)
	ans := []int{tmp.(Num).value}
	heap.Push(h, tmp)
	for i := k; i < len(nums); i++ {
		heap.Push(h, Num{
			index: i,
			value: nums[i],
		})
		tmp := heap.Pop(h).(Num)
		//fmt.Println(tmp)
		for tmp.index <= i-k {
			tmp = heap.Pop(h).(Num)
		}
		ans = append(ans, tmp.value)
		heap.Push(h, tmp)
	}
	return ans
}
