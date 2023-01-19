package k_th_symbol_in_grammar

func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}
	if k > pow(n-2) {
		return reverse(kthGrammar(n-1, k-pow(n-2)))
	} else {
		return kthGrammar(n-1, k)
	}
}

func reverse(k int) int {
	if k == 0 {
		return 1
	} else {
		return 0
	}
}

func pow(n int) (ans int) {
	ans = 1
	for i := 0; i < n; i++ {
		ans *= 2
	}
	return
}
