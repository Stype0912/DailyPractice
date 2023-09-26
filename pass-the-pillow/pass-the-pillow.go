package pass_the_pillow

func passThePillow(n int, time int) int {
	direction := 1
	ans := 1
	for i := 0; i < time; i++ {
		if ans == 1 {
			direction = 1
		} else if ans == n {
			direction = -1
		}
		ans += direction
	}
	return ans
}
