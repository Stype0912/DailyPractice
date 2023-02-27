package array_partition

import "sort"

type IntList []int

func (m IntList) Len() int {
	return len(m)
}

func (m IntList) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m IntList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func arrayPairSum(nums []int) (res int) {
	res = 0
	sort.Sort(IntList(nums))
	for i, num := range nums {
		if i%2 == 1 {
			res += num
		}
	}
	return
}
