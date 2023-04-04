package minimum_flips_to_make_a_or_b_equal_to_c

import (
	"strconv"
)

func minFlips(a int, b int, c int) int {
	aString := strconv.FormatInt(int64(a), 2)
	bString := strconv.FormatInt(int64(b), 2)
	cString := strconv.FormatInt(int64(c), 2)
	maxLength := max(len(aString), max(len(bString), len(cString)))
	tmp := maxLength - len(aString)
	for i := 0; i < tmp; i++ {
		aString = "0" + aString
	}
	tmp = maxLength - len(bString)
	for i := 0; i < tmp; i++ {
		bString = "0" + bString
	}
	tmp = maxLength - len(cString)
	for i := 0; i < tmp; i++ {
		cString = "0" + cString
	}
	ans := 0
	for i := 0; i < len(cString); i++ {
		if cString[i] == '0' {
			if aString[i] == '1' {
				ans++
			}
			if bString[i] == '1' {
				ans++
			}
		} else if cString[i] == '1' {
			if aString[i] == '0' && bString[i] == '0' {
				ans++
			}
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
