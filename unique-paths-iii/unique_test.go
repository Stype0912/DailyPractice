package unique_paths_iii

import "testing"

func TestUniquePath(t *testing.T) {
	t.Log(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
}
