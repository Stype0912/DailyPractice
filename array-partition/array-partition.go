package array_partition

import "sort"

type intList []int

func (m intList) Len() int {
	return len(m)
}

func (m intList) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func arrayPairSum(nums []int) (res int) {
	res = 0
	sort.Sort(intList(nums))
	for i, num := range nums {
		if i%2 == 1 {
			res += num
		}
	}
	return
}
