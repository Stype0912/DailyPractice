package capitalize_the_title

import "strings"

func Capitalize(title string) string {
	ans := ""
	for i := 0; i < len(title); i++ {
		if i == 0 && len(title) >= 3 {
			if title[i] >= 'A' && title[i] <= 'Z' {
				ans += string(title[i])
			} else {
				ans += string('A' + title[i] - 'a')
			}
		} else {
			if title[i] >= 'A' && title[i] <= 'Z' {
				ans += string('a' + title[i] - 'A')
			} else {
				ans += string(title[i])
			}
		}
	}
	return ans
}

func capitalizeTitle(title string) string {
	titles := strings.Split(title, " ")
	for i := 0; i < len(titles); i++ {
		titles[i] = Capitalize(titles[i])
	}
	return strings.Join(titles, " ")
}
