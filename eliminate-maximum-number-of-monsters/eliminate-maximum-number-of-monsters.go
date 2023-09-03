package eliminate_maximum_number_of_monsters

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	arrivalTimes := make([]int, n)
	for i := 0; i < n; i++ {
		arrivalTimes[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(arrivalTimes)
	for i := 0; i < n; i++ {
		if arrivalTimes[i] <= i {
			return i
		}
	}
	return n
}
