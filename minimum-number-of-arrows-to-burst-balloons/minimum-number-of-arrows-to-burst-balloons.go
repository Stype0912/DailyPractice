package minimum_number_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	n := len(points)
	ans := 1

	if n == 0 {
		return ans
	}
	maxRight := points[0][1]
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}
