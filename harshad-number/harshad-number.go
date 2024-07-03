package harshad_number

func sumOfTheDigitsOfHarshadNumber(x int) int {
	ans := 0
	raw := x
	for x != 0 {
		num := x % 10
		x /= 10
		ans += num
	}
	if raw%ans == 0 {
		return ans
	}
	return -1
}
