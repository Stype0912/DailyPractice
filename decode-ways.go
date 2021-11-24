package main

import (
	"strconv"
)

// 0的处理逻辑太复杂了
func numDecodings(s string) int {
	var ans = []int{0}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		if s == "0" {
			return 0
		}
		return 1
	} else if len(s) >= 2 {
		ans[0] = 1
		num, _ := strconv.ParseInt(s[:2], 10, 64)
		if num <= 26 && num != 10 && num != 20 {
			ans = append(ans, 2)
		} else if num == 10 || num == 20 {
			ans = append(ans, 1)
		} else if s[1] == '0' {
			ans = append(ans, 0)
		} else {
			ans = append(ans, 1)
		}
	}
	if len(s) == 2 {
		return ans[1]
	}
	for i := 2; i < len(s); i++ {
		temp := 0
		num, _ := strconv.ParseInt(s[i-1:i+1], 10, 64)
		if num <= 26 && s[i-1] != '0' {
			temp += ans[i-2]
		}
		if s[i] != '0' {
			temp += ans[i-1]
		}
		ans = append(ans, temp)
	}
	return ans[len(ans)-1]
}
