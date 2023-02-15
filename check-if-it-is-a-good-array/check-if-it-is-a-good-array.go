package check_if_it_is_a_good_array

func isGoodArray(nums []int) bool {
	g := 0
	for _, num := range nums {
		g = gcd(g, num)
	}
	if g == 1 {
		return true
	}
	return false
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
