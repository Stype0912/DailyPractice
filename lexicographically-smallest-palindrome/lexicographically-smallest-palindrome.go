package lexicographically_smallest_palindrome

func makeSmallestPalindrome(s string) string {
	min := func(i, j uint8) uint8 {
		if i < j {
			return i
		}
		return j
	}
	n := len(s)
	res := []uint8{}
	for i := 0; i < n; i++ {
		res = append(res, s[i])
	}

	for i := 0; i < n/2; i++ {
		if res[i] == res[n-i-1] {
			continue
		}
		tmp := min(res[i], res[n-i-1])
		res[i], res[n-i-1] = tmp, tmp
	}
	ans := ""
	for i := 0; i < n; i++ {
		ans += string(res[i])
	}
	return ans
}
