package camelcase_matching

func camelMatch(queries []string, pattern string) []bool {
	ans := []bool{}
	for _, query := range queries {
		ans = append(ans, match(query, pattern))
	}
	return ans
}

func match(query, pattern string) bool {
	i, j := 0, 0
	for {
		if i == len(query) && j == len(pattern) {
			return true
		} else if i == len(query) && j != len(pattern) {
			return false
		}
		if j < len(pattern) && query[i] == pattern[j] {
			i++
			j++
		} else if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		} else {
			i++
		}
	}
}
