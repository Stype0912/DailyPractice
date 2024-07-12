package length_of_last_word

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	sSlice := strings.Split(s, "")
	return len(sSlice[len(sSlice)-1])
}
