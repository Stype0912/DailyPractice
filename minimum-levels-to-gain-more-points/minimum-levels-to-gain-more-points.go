package minimum_levels_to_gain_more_points

func minimumLevels(possible []int) int {
	n := len(possible)
	preSum := make([]int, n)
	preSum[0] = func() int {
		if possible[0] == 0 {
			return -1
		} else {
			return 1
		}
	}()
	for i := 1; i < n; i++ {
		if possible[i] == 0 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
	}
	for i := 0; i < n-1; i++ {
		aliceScore := preSum[i]
		bobScore := preSum[n-1] - preSum[i]
		if aliceScore > bobScore {
			return i + 1
		}
	}
	return -1
}
