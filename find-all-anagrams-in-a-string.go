package main

import (
	"reflect"
)

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	width := len(s) - len(p) + 1
	subWidth := len(p)
	ans := make([]int, 0)
	window := make(map[string]int)
	subWindow := make(map[string]int)
	letter := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26; i++ {
		window[string(letter[i])]=0
		subWindow[string(letter[i])]=0
	}
	for i := 0; i < subWidth; i++ {
		window[string(s[i])]++
		subWindow[string(p[i])]++
	}
	if reflect.DeepEqual(window, subWindow) {
		ans = append(ans, 0)
	}
	for i := 1; i < width; i++ {
		window[string(s[i-1])]--
		window[string(s[i+subWidth-1])]++
		if reflect.DeepEqual(window, subWindow) {
			ans = append(ans, i)
		}
	}
	return ans
}