package minimum_interval_to_include_each_query

import "testing"

func TestInterval(t *testing.T) {
	t.Log(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}))
}

//[[2,3],[2,5],[1,8],[20,25]]
//[2,19,5,22]
