package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	var res []string
	hash := make(map[string]bool, 0)
	for _, i := range generateParenthesis(n - 1) {
		for j := 0; j < len(i); j++ {
			if !hash[(i[:j] + "()" + i[j:])] {
				res = append(res, i[:j]+"()"+i[j:])
				hash[(i[:j] + "()" + i[j:])] = true
			}
		}
	}
	return res
}
