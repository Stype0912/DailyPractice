package maximize_distance_to_closest_person

func maxDistToClosest(seats []int) int {
	seatMap := []int{}
	for i, seat := range seats {
		if seat == 1 {
			seatMap = append(seatMap, i)
		}
	}
	ans := 1
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if seats[0] != 1 {
		ans = max(ans, seatMap[0])
	}
	if seats[len(seats)-1] != 1 {
		ans = max(ans, len(seats)-1-seatMap[len(seatMap)-1])
	}
	if len(seatMap) == 1 {
		return ans
	}
	for i := 0; i < len(seatMap)-1; i++ {
		ans = max(ans, (seatMap[i+1]-seatMap[i])/2)
	}
	return ans
}
