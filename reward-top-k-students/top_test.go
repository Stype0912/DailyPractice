package reward_top_k_students

import "testing"

func TestTop(t *testing.T) {
	t.Log(topStudents([]string{"smart", "brilliant", "studious"}, []string{"not"}, []string{"this student is not studious", "the student is smart"}, []int{1, 2}, 2))
}
