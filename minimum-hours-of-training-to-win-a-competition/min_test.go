package minimum_hours_of_training_to_win_a_competition

import "testing"

func TestMin(t *testing.T) {
	t.Log(minNumberOfHours(1, 1, []int{1, 1, 1, 1}, []int{1, 1, 1, 50}))
}
