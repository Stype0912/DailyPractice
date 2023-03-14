package target_sum

import "testing"

func TestTarget(t *testing.T) {
	t.Log(findTargetSumWays([]int{2, 107, 109, 113, 127, 131, 137, 3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 47, 53}, 1000))
}
