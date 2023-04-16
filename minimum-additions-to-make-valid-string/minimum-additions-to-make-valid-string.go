package minimum_additions_to_make_valid_string

func addMinimum(word string) int {
	ans := 0
	for i := 0; i < len(word); {
		if word[i] == 'a' {
			if i < len(word)-1 && word[i+1] == 'b' {
				if i < len(word)-2 && word[i+2] == 'c' {
					i += 3
				} else {
					ans++
					i += 2
				}
			} else if i < len(word)-1 && word[i+1] == 'c' {
				ans += 1
				i += 2
			} else {
				ans += 2
				i += 1
			}
			continue
		} else if word[i] == 'b' {
			ans++
			if i < len(word)-1 && word[i+1] == 'c' {
				i += 2
			} else {
				ans++
				i += 1
			}
			continue
		} else if word[i] == 'c' {
			i += 1
			ans += 2
		}
	}
	return ans
}
