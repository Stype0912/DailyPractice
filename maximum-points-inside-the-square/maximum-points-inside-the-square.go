package maximum_points_inside_the_square

import (
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func maxPointsInsideSquare(points [][]int, s string) int {
	n := len(points)
	if n == 1 {
		return n
	}
	squareLength := make([][]int, n)
	for i, point := range points {
		length := max(abs(point[0]), abs(point[1]))
		squareLength[i] = []int{length, i}
	}
	sort.Slice(squareLength, func(i, j int) bool {
		return squareLength[i][0] < squareLength[j][0]
	})
	tagMap := map[uint8]int{}
	ans := 0
	tagMap[s[squareLength[0][1]]] = 1
	for i := 1; i < n; i++ {
		if tagMap[s[squareLength[i][1]]] == 0 {
			tagMap[s[squareLength[i][1]]] = 1
			if squareLength[i][0] != squareLength[i-1][0] {
				ans = i
			}
			if squareLength[i][0] == squareLength[i-1][0] && i == n-1 {
				ans = i
			}
		} else {
			if squareLength[i][0] != squareLength[i-1][0] {
				ans = i
			}
			return ans
		}
	}
	return ans
}
