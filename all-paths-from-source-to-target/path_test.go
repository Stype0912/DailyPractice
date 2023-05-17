package all_paths_from_source_to_target

import "testing"

func TestPath(t *testing.T) {
	t.Log(allPathsSourceTarget([][]int{{1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}))
}
