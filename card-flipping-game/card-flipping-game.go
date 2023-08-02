package card_flipping_game

import "math"

func flipgame(fronts []int, backs []int) int {
	numMap := map[int]bool{}
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			numMap[fronts[i]] = true
		}
	}
	ans := math.MaxInt
	for i := 0; i < len(fronts); i++ {
		if !numMap[fronts[i]] && fronts[i] < ans {
			ans = fronts[i]
		}
		if !numMap[backs[i]] && backs[i] < ans {
			ans = backs[i]
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
