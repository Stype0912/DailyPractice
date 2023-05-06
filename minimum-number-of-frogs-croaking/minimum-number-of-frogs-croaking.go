package minimum_number_of_frogs_croaking

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}
	frogMap := map[byte]int{}
	for i := 0; i < len(croakOfFrogs); i++ {
		switch croakOfFrogs[i] {
		case 'c':
			if frogMap['k'] != 0 {
				frogMap['k']--
				frogMap['c']++
			} else {
				frogMap['c']++
			}
		case 'r':
			if frogMap['c'] != 0 {
				frogMap['c']--
				frogMap['r']++
			} else {
				return -1
			}
		case 'o':
			if frogMap['r'] != 0 {
				frogMap['r']--
				frogMap['o']++
			} else {
				return -1
			}
		case 'a':
			if frogMap['o'] != 0 {
				frogMap['o']--
				frogMap['a']++
			} else {
				return -1
			}
		case 'k':
			if frogMap['a'] != 0 {
				frogMap['a']--
				frogMap['k']++
			} else {
				return -1
			}
		}
	}
	for _, letter := range []byte{'c', 'r', 'o', 'a'} {
		if frogMap[letter] != 0 {
			return -1
		}
	}
	return frogMap['k']
}
