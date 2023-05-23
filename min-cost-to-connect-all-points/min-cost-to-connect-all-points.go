package min_cost_to_connect_all_points

import "sort"

func minCostConnectPoints(points [][]int) int {
	edges := [][]int{}
	pointsColor := map[int]int{}
	ans := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, []int{absMinus(points[i][0], points[j][0]) + absMinus(points[i][1], points[j][1]), i, j})
		}
		pointsColor[i] = i + 1
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	for _, edge := range edges {
		if pointsColor[edge[1]] != pointsColor[edge[2]] {
			ans += edge[0]
			tmpMin := min(pointsColor[edge[1]], pointsColor[edge[2]])
			tmpMax := pointsColor[edge[1]] + pointsColor[edge[2]] - tmpMin
			for key, value := range pointsColor {
				if value == tmpMax {
					pointsColor[key] = tmpMin
				}
			}
			//pointsColor[edge[1]] = pointsColor[edge[2]]
		}
	}
	return ans
}

func absMinus(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
