package WHnhjV

import "math"

func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		tmp := gem[operation[0]] / 2
		gem[operation[0]] -= tmp
		gem[operation[1]] += tmp
	}
	minGem := math.MaxInt
	maxGem := math.MinInt
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(gem); i++ {
		minGem = min(minGem, gem[i])
		maxGem = max(maxGem, gem[i])
	}
	return maxGem - minGem
}
