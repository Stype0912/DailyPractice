package course_schedule_ii

import "testing"

func TestCourse(t *testing.T) {
	t.Log(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
