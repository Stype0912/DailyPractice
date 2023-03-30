package widest_vertical_area_between_two_points_containing_no_points

import "sort"

type Points [][]int

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Sort(Points(points))
	ans := 0
	for i := 1; i < len(points); i++ {
		ans = max(ans, points[i][0]-points[i-1][0])
	}
	return ans
}
