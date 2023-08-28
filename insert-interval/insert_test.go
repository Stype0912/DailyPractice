package insert_interval

import "testing"

func TestInsert(t *testing.T) {
	t.Log(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
