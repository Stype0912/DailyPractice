package max_value_of_equation

import "math"

func findMaxValueOfEquation(points [][]int, k int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	ans := math.MinInt
	type pair struct {
		x  int
		yx int
	}
	q := []pair{}
	for _, p := range points {
		x, y := p[0], p[1]
		for len(q) > 0 && q[0].x < x-k {
			q = q[1:]
		}
		if len(q) > 0 {
			ans = max(ans, x+y+q[0].yx)
		}
		for len(q) > 0 && q[len(q)-1].yx <= y-x {
			q = q[:len(q)-1]
		}
		q = append(q, pair{x, y - x})
	}
	return ans
}
