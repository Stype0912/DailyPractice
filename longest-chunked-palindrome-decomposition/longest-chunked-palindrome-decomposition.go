package longest_chunked_palindrome_decomposition

func longestDecomposition(text string) int {
	if len(text) == 0 {
		return 0
	}
	ans := 1
	for i := 0; i < len(text); i++ {
		j := len(text) - i - 1
		if i >= j {
			break
		}
		if text[:i+1] == text[j:] {
			return 2 + longestDecomposition(text[i+1:j])
		}
	}
	return ans
}
