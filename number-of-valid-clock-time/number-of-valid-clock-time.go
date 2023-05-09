package number_of_valid_clock_time

func countTime(time string) int {
	ans := 1
	switch time[0] {
	case '?':
		if time[1] == '?' {
			ans = 24
		} else if time[1] <= '3' {
			ans = 3
		} else {
			ans = 2
		}
	case '0', '1':
		if time[1] == '?' {
			ans = 10
		}
	case '2':
		if time[1] == '?' {
			ans = 4
		}
	}
	switch time[3] {
	case '?':
		if time[4] == '?' {
			ans *= 60
		} else {
			ans *= 6
		}
	default:
		if time[4] == '?' {
			ans *= 10
		}
	}
	return ans
}
