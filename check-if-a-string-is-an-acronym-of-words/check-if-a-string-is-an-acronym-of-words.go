package check_if_a_string_is_an_acronym_of_words

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	ans := ""
	for _, word := range words {
		ans += string(word[0])
	}
	return ans == s
}
