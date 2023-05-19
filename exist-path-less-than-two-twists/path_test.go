package exist_path_less_than_two_twists

import "testing"

func TestPath(t *testing.T) {
	t.Log(pathExists([][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 2, 2, 2, 0},
		{0, 0, 0, 2, 2, 0},
		{0, 2, 0, 0, 2, 0},
		{0, 0, 1, 2, 2, 0},
		{0, 0, 0, 0, 0, 0}}))
}
