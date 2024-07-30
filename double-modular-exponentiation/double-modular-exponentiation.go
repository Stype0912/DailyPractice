package double_modular_exponentiation

func getGoodIndices(variables [][]int, target int) []int {
	pow_mod := func(a, b, mod int) int {
		ans := 1
		for b > 0 {
			if b%2 == 1 {
				ans = ans * a % mod
			}
			a = a * a % mod
			b >>= 1
		}
		return ans
	}
	ans := []int{}
	for i, v := range variables {
		if pow_mod(pow_mod(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return ans
}
