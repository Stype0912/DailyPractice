package beautiful_towers_ii

func maximumSumOfHeights(maxHeights []int) int64 {
	stackLeft := []int{}
	prefix := make([]int, len(maxHeights))
	for i := 0; i < len(maxHeights); i++ {
		for len(stackLeft) > 0 && maxHeights[i] < maxHeights[stackLeft[len(stackLeft)-1]] {
			stackLeft = stackLeft[:len(stackLeft)-1]
		}
		if len(stackLeft) == 0 {
			prefix[i] = maxHeights[i] * (i + 1)
		} else {
			last := stackLeft[len(stackLeft)-1]
			prefix[i] = prefix[last] + maxHeights[i]*(i-last)
		}
		stackLeft = append(stackLeft, i)
	}
	stackRight := []int{}
	suffix := make([]int, len(maxHeights))
	res := int64(0)
	max := func(i, j int64) int64 {
		if i > j {
			return i
		}
		return j
	}
	for i := len(maxHeights) - 1; i >= 0; i-- {
		for len(stackRight) > 0 && maxHeights[i] < maxHeights[stackRight[len(stackRight)-1]] {
			stackRight = stackRight[:len(stackRight)-1]
		}
		if len(stackRight) == 0 {
			suffix[i] = maxHeights[i] * (len(maxHeights) - i)
		} else {
			last := stackRight[len(stackRight)-1]
			suffix[i] = suffix[last] + maxHeights[i]*(last-i)
		}
		stackRight = append(stackRight, i)
		res = max(res, int64(prefix[i])+int64(suffix[i])-int64(maxHeights[i]))
	}
	return res
}
