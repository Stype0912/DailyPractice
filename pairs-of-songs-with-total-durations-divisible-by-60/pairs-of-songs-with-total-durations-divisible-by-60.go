package pairs_of_songs_with_total_durations_divisible_by_60

func numPairsDivisibleBy60(time []int) int {
	timeModulesMap := map[int]int{}
	for _, timeInt := range time {
		timeModulesMap[timeInt%60]++
	}
	ans := 0
	for key, value := range timeModulesMap {
		if key == 0 || key == 30 {
			ans += value * (value - 1)
		} else {
			ans += timeModulesMap[key] * timeModulesMap[60-key]
		}
	}
	return ans / 2
}
