package minimum_number_of_taps_to_open_to_water_a_garden

import "sort"

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

type doubleInt [][]int

func (d doubleInt) Len() int {
	return len(d)
}
func (d doubleInt) Less(i int, j int) bool {
	if d[i][0] == d[j][0] {
		return d[i][1] < d[j][1]
	} else {
		return d[i][0] < d[j][0]
	}
}
func (d doubleInt) Swap(i int, j int) {
	d[i], d[j] = d[j], d[i]
}

func minTaps1(n int, ranges []int) int {
	ans := [][]int{}
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
	areas := [][]int{}
	for i := 0; i <= n; i++ {
		l := max(i-ranges[i], 0)
		r := min(i+ranges[i], n)
		if l != r {
			areas = append(areas, []int{l, r})
		}
	}
	sort.Sort(doubleInt(areas))
	for _, area := range areas {
		if len(ans) == 0 {
			ans = append(ans, area)
			continue
		}
		last := ans[len(ans)-1]
		oldLeft, oldRight := last[0], last[1]
		newLeft, newRight := area[0], area[1]
		if newLeft > oldRight {
			return -1
		}
		if newLeft == oldLeft {
			ans = ans[:len(ans)-1]
			ans = append(ans, area)
			continue
		}
		if newLeft == oldRight {
			ans = append(ans, area)
			continue
		}
		if newRight <= oldRight {
			continue
		}
		for len(ans) >= 2 {
			lastLast := ans[len(ans)-2]
			if newLeft <= lastLast[1] {
				ans = ans[:len(ans)-1]
			} else {
				break
			}
		}
		ans = append(ans, area)
	}
	if ans[0][0] > 0 || ans[len(ans)-1][1] < n || len(ans) == 0 {
		return -1
	} else {
		return len(ans)
	}
}
