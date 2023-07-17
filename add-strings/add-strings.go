package add_strings

func addStrings(num1 string, num2 string) string {
	lenNum1, lenNum2 := len(num1), len(num2)
	maxLen := max(lenNum1, lenNum2)
	flag := 0
	ansTmp := []int{}
	ans := ""
	for i := 0; i < maxLen; i++ {
		first := lenNum1 - i - 1
		second := lenNum2 - i - 1
		sum := 0
		if first >= 0 {
			sum += int(num1[first]) - '0'
		}
		if second >= 0 {
			sum += int(num2[second]) - '0'
		}
		ansTmp = append(ansTmp, '0'+((sum+flag)%10))
		if sum+flag > 9 {
			flag = 1
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		ansTmp = append(ansTmp, '0'+1)
	}
	for i := len(ansTmp) - 1; i >= 0; i-- {
		ans += string(rune(ansTmp[i]))
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
