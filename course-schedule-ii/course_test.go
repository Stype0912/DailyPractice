package course_schedule_ii

import "testing"

func TestCourse(t *testing.T) {
	t.Log(findOrder(2, [][]int{{1, 0}, {0, 1}}))
}
