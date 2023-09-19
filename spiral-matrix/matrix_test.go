package spiral_matrix

import "testing"

func TestMatrix(t *testing.T) {
	t.Log(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
