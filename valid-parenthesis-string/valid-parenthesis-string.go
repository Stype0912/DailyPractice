package valid_parenthesis_string

func checkValidString(s string) bool {
	stackBracket := []int{}
	stackStar := []int{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stackBracket = append(stackBracket, i)
		case '*':
			stackStar = append(stackStar, i)
		case ')':
			if len(stackBracket) == 0 && len(stackStar) == 0 {
				return false
			}
			if len(stackBracket) != 0 {
				stackBracket = stackBracket[:len(stackBracket)-1]
			} else if len(stackStar) != 0 {
				stackStar = stackStar[:len(stackStar)-1]
			}
		}
	}
	if len(stackBracket) == 0 {
		return true
	} else if len(stackStar) > len(stackBracket) {
		n := len(stackBracket)
		for i := 0; i < n; i++ {
			if stackStar[len(stackStar)-1] > stackBracket[len(stackBracket)-1] {
				stackBracket = stackBracket[:len(stackBracket)-1]
				stackStar = stackStar[:len(stackStar)-1]
				continue
			} else {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
