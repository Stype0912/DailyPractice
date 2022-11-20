package minimum_number_of_swaps_to_make_the_string_balanced

func minSwaps(s string) int {
	res := 0
	counter := 0
	for _, item := range s {
		if item == '[' {
			counter += 1
		} else {
			counter -= 1
			if counter < res {
				res = counter
			}
		}
	}
	res = ((-1)*res + 1) / 2
	return res
}
