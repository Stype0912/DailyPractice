package maximum_number_of_consecutive_values_you_can_make

import "sort"

type intList []int

func (m intList) Len() int {
	return len(m)
}

func (m intList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func getMaximumConsecutive(coins []int) int {
	switch len(coins) {
	case 0:
		return 1
	default:
		sort.Sort(intList(coins))
		res := 1
		for _, coin := range coins {
			if res+1 >= coin {
				res += coin
			} else {
				break
			}
		}
		return res
	}
}
