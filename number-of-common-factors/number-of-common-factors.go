package number_of_common_factors

func commonFactors(a int, b int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	ans := 0
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			ans++
		}
	}
	return ans
}
