package powerful_integers

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	ansMap := map[int]bool{}
	for i := 0; ; i++ {
		if int(math.Pow(float64(x), float64(i))) > bound {
			break
		}
		for j := 0; ; j++ {
			tmp := int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j)))
			if y == 1 {
				if tmp <= bound {
					ansMap[tmp] = true
				}
				break
			}
			if tmp > bound {
				break
			} else {
				ansMap[tmp] = true
			}
		}
		if x == 1 {
			break
		}
	}
	ans := []int{}
	for key, _ := range ansMap {
		ans = append(ans, key)
	}
	return ans
}
