package maximum_number_of_balls_in_a_box

func countBalls(lowLimit int, highLimit int) int {
	boxCountMap := make(map[int]int)
	maxRes := 0
	for i := lowLimit; i <= highLimit; i++ {
		boxCountMap[sumOfNum(i)]++
		if boxCountMap[sumOfNum(i)] > maxRes {
			maxRes = boxCountMap[sumOfNum(i)]
		}
	}
	return maxRes
}

func sumOfNum(num int) int {
	res := 0
	for {
		if num == 0 {
			return res
		}
		res += num % 10
		num = num / 10
	}
}
