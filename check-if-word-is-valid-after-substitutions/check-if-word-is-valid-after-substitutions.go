package check_if_word_is_valid_after_substitutions

func isValid(s string) bool {
	flag := false
	for {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == 'a' && s[i+1] == 'b' && s[i+2] == 'c' {
				s = s[:i] + s[i+3:]
				flag = true
				break
			}
		}
		if len(s) != 0 && !flag {
			return false
		}
		if len(s) == 0 {
			return true
		}
		flag = false
	}
}
