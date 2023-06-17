package minimum_cuts_to_divide_a_circle

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	} else {
		return n
	}
}
