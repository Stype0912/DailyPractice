package pairs_of_songs_with_total_durations_divisible_by_60

import "testing"

func TestPair(t *testing.T) {
	t.Log(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
}
