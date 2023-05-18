package dijkstra

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	var gp = [][]int{
		{0, 100, 1200, INF, INF, INF},
		{100, 0, 900, 300, INF, INF},
		{1200, 900, 0, 400, 500, INF},
		{INF, 300, 400, 0, 1300, 1400},
		{INF, INF, 500, 1300, 0, 1500},
		{INF, INF, INF, 1400, 1500, 0},
	}
	dist := dijkstra(gp, 0)
	t.Log(dist)
}
