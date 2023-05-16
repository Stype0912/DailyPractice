package minimum_difficulty_of_a_job_schedule

import "testing"

func TestJob(t *testing.T) {
	t.Log(minDifficulty([]int{7, 1, 7, 1, 7, 1}, 3))
}
