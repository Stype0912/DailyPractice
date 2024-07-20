package reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	sList := strings.Fields(s)
	n := len(sList)
	for i := 0; i < n/2; i++ {
		sList[i], sList[n-1-i] = sList[n-1-i], sList[i]
	}
	return strings.Join(sList, " ")
}
