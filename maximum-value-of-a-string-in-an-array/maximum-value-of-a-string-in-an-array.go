package maximum_value_of_a_string_in_an_array

import "strconv"

func maximumValue(strs []string) int {
	ans := 0
	for _, str := range strs {
		num, err := strconv.ParseInt(str, 10, 64)
		tmp := 0
		if err != nil {
			tmp = len(str)
		} else {
			tmp = int(num)
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
