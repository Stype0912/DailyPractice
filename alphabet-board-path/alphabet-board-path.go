package alphabet_board_path

func alphabetBoardPath(target string) string {
	now := 'a'
	var row, line int
	res := ""
	for _, letter := range target {
		if letter == 'z' {
			if now == 'z' {
				res += "!"

			} else {
				letter = 'u'
				row = int((letter-'a')/5) - int((now-'a')/5)
				line = int((letter-'a')%5) - int((now-'a')%5)

				if row > 0 {
					for k := 0; k < row; k++ {
						res += "D"
					}
				} else if row < 0 {
					for k := 0; k > row; k-- {
						res += "U"
					}
				}

				if line > 0 {
					for k := 0; k < line; k++ {
						res += "R"
					}
				} else if line < 0 {
					for k := 0; k > line; k-- {
						res += "L"
					}
				}
				res += "D!"
				now = 'z'
			}
		} else {
			if now == 'z' {
				res += "U"
				now = 'u'
			}
			row = int((letter-'a')/5) - int((now-'a')/5)
			line = int((letter-'a')%5) - int((now-'a')%5)

			if row > 0 {
				for k := 0; k < row; k++ {
					res += "D"
				}
			} else if row < 0 {
				for k := 0; k > row; k-- {
					res += "U"
				}
			}

			if line > 0 {
				for k := 0; k < line; k++ {
					res += "R"
				}
			} else if line < 0 {
				for k := 0; k > line; k-- {
					res += "L"
				}
			}
			res += "!"
			now = letter
		}

	}
	return res
}
