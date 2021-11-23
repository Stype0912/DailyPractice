package main

func buddyStrings(s string, goal string) bool {
	if s == goal {
		HashMap := make(map[string]int)
		for _, item := range s {
			HashMap[string(item)]++
		}
		for _, item := range HashMap {
			if item >= 2 {
				return true
			}
		}
		return false
	}
	length := len(s)
	if length != len(goal) {
		return false
	}
	count := 0
	var sa, sb, ga, gb string
	for i := 0; i < length; i++ {
		if s[i] != goal[i] {
			count++
			if count == 1 {
				sa = string(s[i])
				ga = string(goal[i])
			} else if count == 2 {
				sb = string(s[i])
				gb = string(goal[i])
			}
		}
	}
	if count != 2 {
		return false
	}
	if sa == gb && sb == ga {
		return true
	}
	return false
}
