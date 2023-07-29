package delete_greatest_value_in_each_row

import "testing"

func TestDelete(t *testing.T) {
	t.Log(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
}
