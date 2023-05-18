package minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	TwordMap := map[byte]int{}
	SwordMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		TwordMap[t[i]]++
	}
	left := 0
	ans := ""
	minLength := math.MaxInt
	for right := 0; right < len(s); right++ {
		SwordMap[s[right]]++
		for isValid(SwordMap, TwordMap) {
			if isValid(SwordMap, TwordMap) {
				if len(s[left:right+1]) < minLength {
					ans = s[left : right+1]
					minLength = right - left + 1
				}
			}
			SwordMap[s[left]]--
			left++
		}
	}
	return ans
}

func isValid(swordMap, twordMap map[byte]int) bool {
	for key, value := range twordMap {
		if swordMap[key] < value {
			return false
		}
	}
	return true
}
