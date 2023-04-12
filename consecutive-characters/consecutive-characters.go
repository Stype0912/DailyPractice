package consecutive_characters

func maxPower(s string) int {
	ans := 1
	tmp := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			tmp++
		} else {
			tmp = 1
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
