package slowest_key

func slowestKey(releaseTimes []int, keysPressed string) byte {
	releaseTimes = append([]int{0}, releaseTimes...)
	var ansByte byte
	var longestTime int = 0
	for i := 0; i < len(keysPressed); i++ {
		lastTime := releaseTimes[i+1] - releaseTimes[i]
		if lastTime > longestTime || (lastTime == longestTime && keysPressed[i] > ansByte) {
			longestTime = lastTime
			ansByte = keysPressed[i]
		}
	}
	return ansByte
}
