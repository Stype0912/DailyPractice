package alternating_digit_sum

import "strconv"

func alternateDigitSum(n int) int {
	number := strconv.FormatInt(int64(n), 10)
	ans := 0
	for i := 0; i < len(number); i++ {
		tmp, _ := strconv.ParseInt(string(int(number[i])), 10, 64)
		if i%2 == 0 {
			ans += int(tmp)
		} else {
			ans -= int(tmp)
		}
	}
	return ans
}
