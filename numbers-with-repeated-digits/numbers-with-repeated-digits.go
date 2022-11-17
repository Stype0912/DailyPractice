package main

import (
	"strconv"
)

func numDupDigitsAtMostN(n int) int {
	if n <= 10 {
		return 0
	}
	ans := 0
	return ans
}

func jiecheng(num int) int {
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}

func contain(item int) bool {
	var HashMap = map[rune]int{
		'0': 0,
		'1': 0,
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
	}
	s := strconv.FormatInt(int64(item), 10)
	for _, item := range s {
		if HashMap[item] != 0 {
			return true
		}
		HashMap[item] += 1
	}
	return false
}
