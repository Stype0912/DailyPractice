package minimum_garden_perimeter_to_collect_enough_apples

func minimumPerimeter(neededApples int64) int64 {
	apples := int64(0)
	i := 1
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	for apples < neededApples {
		tmp := 8 * i
		for j := -i + 1; j < i; j++ {
			tmp += 4 * (i + abs(j))
		}
		apples += int64(tmp)
		i++
	}
	return int64(8 * (i - 1))
}
