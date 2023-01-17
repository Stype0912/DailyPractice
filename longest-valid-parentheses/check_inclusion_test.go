package longest_valid_parentheses

import "testing"

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range s1 {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0

	if len(s1) > len(s2) {
		return false
	}

	for right < len(s2) {
		tempAdd := s2[right]
		right++
		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}
		if right-left == len(s1) {
			if valid == len(need) {
				return true
			}
			tempDel := s2[left]
			left++
			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] == need[tempDel] {
					valid--
				}
			}
		}
	}
	return false
}

func TestCheck(t *testing.T) {
	t.Log(checkInclusion("aa", "aaa"))
}
