package main

import (
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	length := len(encodedText)
	if length == 0 {
		return ""
	}
	if rows == 1 {
		return strings.TrimRight(encodedText, " ")
	}
	lines := length / rows
	ans := ""
	i, j := 0, 0
	head := j
	for {
		ans += string(encodedText[lines*i+j])
		i++
		j++
		if i == rows {
			i = 0
			j = head + 1
			head++
		}
		if j == lines {
			break
		}
	}
	return strings.TrimRight(ans, " ")
}