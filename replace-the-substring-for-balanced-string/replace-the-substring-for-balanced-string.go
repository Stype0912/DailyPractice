package replace_the_substring_for_balanced_string

func balancedString(s string) int {
	count := map[rune]int{}
	for _, letter := range s {
		count[letter]++
	}
	partial := len(s) / 4
	check := func() bool {
		if count['Q'] > partial ||
			count['W'] > partial ||
			count['E'] > partial ||
			count['R'] > partial {
			return false
		}
		return true
	}
	if check() {
		return 0
	}
	r := 0
	res := len(s)
	for l, letter := range s {
		for r < len(s) && !check() {
			count[rune(s[r])] -= 1
			r += 1
		}
		if !check() {
			return res
		}
		res = func(a, b int) int {
			if a < b {
				return a
			} else {
				return b
			}
		}(res, r-l)
		count[letter]++
	}
	return res
}
