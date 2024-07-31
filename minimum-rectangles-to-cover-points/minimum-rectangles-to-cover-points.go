package minimum_rectangles_to_cover_points

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 0
	bound := -1

	for i := 0; i < len(points); i++ {
		if points[i][0] <= bound {
			continue
		}
		res++
		bound = points[i][0] + w
	}
	return res
}
