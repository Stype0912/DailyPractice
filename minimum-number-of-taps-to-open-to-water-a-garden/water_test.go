package minimum_number_of_taps_to_open_to_water_a_garden

import "testing"

func TestWater(t *testing.T) {
	t.Log(minTaps(4, []int{0, 1, 0, 1, 0}))
}
