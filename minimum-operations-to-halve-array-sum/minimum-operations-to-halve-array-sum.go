package minimum_operations_to_halve_array_sum

import "container/heap"

type NumsInt []float64

func (n NumsInt) Len() int {
	return len(n)
}

func (n NumsInt) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n NumsInt) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NumsInt) Push(x interface{}) {
	*n = append(*n, x.(float64))
}

func (n *NumsInt) Pop() interface{} {
	old := *n
	tmp := len(old)
	x := old[tmp-1]
	*n = old[:tmp-1]
	return x
}

func halveArray(nums []int) int {
	numsInt := &NumsInt{}
	heap.Init(numsInt)
	sum := 0
	for _, num := range nums {
		heap.Push(numsInt, float64(num))
		sum += num
	}
	minus := 0.0
	ans := 0
	for minus < float64(sum)/2.0 {
		largest := heap.Pop(numsInt).(float64)
		minus += largest / 2.0
		heap.Push(numsInt, largest/2.0)
		ans++
	}
	return ans
}
