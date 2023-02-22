package minimum_number_of_taps_to_open_to_water_a_garden

import "testing"

func TestWater(t *testing.T) {
	t.Log(minTaps1(5, []int{3, 4, 1, 1, 0, 0}))
}
