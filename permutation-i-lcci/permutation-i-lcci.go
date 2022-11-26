package permutation_i_lcci

var res [][]rune

func permutation(S string) []string {
	res = [][]rune{}
	var resFinal []string
	dfs([]rune(S[:]), 0)
	for _, item := range res {
		resFinal = append(resFinal, string(item))
	}
	return resFinal
}

func dfs(S []rune, i int) {
	if i == len(S) {
		res = append(res, []rune{})
		res[len(res)-1] = append(res[len(res)-1], S...)
	} else {
		for j := i; j < len(S); j++ {
			S[i], S[j] = S[j], S[i]
			dfs(S, i+1)
			S[i], S[j] = S[j], S[i]
		}
	}
}
