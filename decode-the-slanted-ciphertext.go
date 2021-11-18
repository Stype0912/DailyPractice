package main

import (
	"fmt"
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
	matrix := make([][]string, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]string, lines)
	}
	for i := 0; i < length; i++ {
		matrix[i/lines][i%lines] = string(encodedText[i])
	}
	ans := ""
	i, j := 0, 0
	head := j
	for {
		ans += matrix[i][j]
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

func main() {
	fmt.Println(decodeCiphertext(" b  ac", 2))
}
