package split_two_strings_to_make_palindrome

func checkPalindromeFormation(a string, b string) bool {
	return checkPalindrome(a, b) || checkPalindrome(b, a)
}

func checkPalindrome(a, b string) bool {
	i, j := 0, len(a)-1
	for i < j && a[i] == b[j] {
		i++
		j--
	}
	return isReverse(a[i : j+1])
}

func isReverse(a string) bool {
	i, j := 0, len(a)-1
	for i < j && a[i] == a[j] {
		i++
		j--
	}
	return i >= j
}
