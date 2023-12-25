package number_of_burgers_with_no_waste_of_ingredients

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices&1 == 1 {
		return []int{}
	}
	if tomatoSlices/2-cheeseSlices < 0 || 2*cheeseSlices-tomatoSlices/2 < 0 {
		return []int{}
	}
	return []int{tomatoSlices/2 - cheeseSlices, 2*cheeseSlices - tomatoSlices/2}
}
