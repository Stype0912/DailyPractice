package minimum_number_of_taps_to_open_to_water_a_garden

func minTaps(n int, ranges []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	land := map[int]int{}
	for i := 0; i <= n; i++ {
		l := max(i-ranges[i], 0)
		r := min(i+ranges[i], n)
		for j := l; j < r; j++ {
			if _, ok := land[j]; ok {
				land[j] = max(land[j], r)
			} else {
				land[j] = r
			}
		}
	}
	cnt := 0
	for i := 0; i < n; {
		if j, ok := land[i]; !ok {
			return -1
		} else {
			i = j
			cnt++
		}
	}
	return cnt
}
