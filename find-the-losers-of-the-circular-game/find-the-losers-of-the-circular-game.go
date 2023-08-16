package find_the_losers_of_the_circular_game

func circularGameLosers(n int, k int) []int {
	participantMap := make([]int, n)
	ball := 0
	round := 1
	for {
		if participantMap[ball] != 0 {
			break
		}
		participantMap[ball] = 1
		ball += round * k
		ball %= n
		round++
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		if participantMap[i] == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
