package longest_well_performing_interval

func longestWPI(hours []int) int {
	pos := map[int]int{}
	s, ans := 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i, hour := range hours {
		if hour > 8 {
			s++
		} else {
			s--
		}
		if s > 0 {
			ans = i + 1
		}
		if s <= 0 {
			if j, ok := pos[s-1]; ok {
				ans = max(ans, i-j)
			}
			if _, ok := pos[s]; !ok {
				pos[s] = i
			}
		}
	}
	return ans
}
