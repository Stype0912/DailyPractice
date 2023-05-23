package min_cost_to_connect_all_points

import "testing"

func TestCost(t *testing.T) {
	t.Log(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}
