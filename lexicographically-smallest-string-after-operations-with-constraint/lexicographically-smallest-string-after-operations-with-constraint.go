package lexicographically_smallest_string_after_operations_with_constraint

func getSmallestString(s string, k int) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			ans += "a"
			continue
		}
		if k > 0 {
			distance := int(min('a'-s[i]+26, s[i]-'a'))
			if k >= distance {
				k -= distance
				ans += "a"
			} else {
				word1, word2 := int(s[i])+k, int(s[i])-k
				if word1 > 'z' {
					word1 = word1 - 'z' + 'a'
				}
				if word2 < 'a' {
					word2 = 'z' - ('a' - word2)
				}
				ans += string(byte(min(word1, word2)))
				k = 0
			}
		} else {
			ans += string(s[i])
		}
	}
	return ans
}
