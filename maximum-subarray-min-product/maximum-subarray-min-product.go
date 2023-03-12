package maximum_subarray_min_product

func maxSumMinProduct(nums []int) int {
	max := func(i, j int64) int64 {
		if i > j {
			return i
		} else {
			return j
		}
	}
	rightStack := []int{}
	rightPos := map[int]int{len(nums) - 1: len(nums)}
	for i := len(nums) - 1; i >= 0; i-- {
		if len(rightStack) != 0 {
			for len(rightStack) != 0 && nums[i] <= nums[rightStack[len(rightStack)-1]] {
				rightStack = rightStack[:len(rightStack)-1]
			}
			if len(rightStack) == 0 {
				rightPos[i] = len(nums)
			} else {
				rightPos[i] = rightStack[len(rightStack)-1]
			}
		}
		rightStack = append(rightStack, i)
	}

	leftStack := []int{}
	leftPos := map[int]int{0: -1}
	prefixSum := map[int]int64{-1: 0}
	for i := 0; i < len(nums); i++ {
		if len(leftStack) != 0 {
			for len(leftStack) != 0 && nums[i] <= nums[leftStack[len(leftStack)-1]] {
				leftStack = leftStack[:len(leftStack)-1]
			}
			if len(leftStack) == 0 {
				leftPos[i] = -1
			} else {
				leftPos[i] = leftStack[len(leftStack)-1]
			}
		}
		leftStack = append(leftStack, i)
		prefixSum[i] = prefixSum[i-1] + int64(nums[i])
	}
	ans := int64(0)
	for i, num := range nums {
		leftBorder := leftPos[i]
		rightBorder := rightPos[i]
		ans = max(ans, int64(num)*(prefixSum[rightBorder-1]-prefixSum[leftBorder]))
		//fmt.Println(i, num, leftBorder, rightBorder, ans)
	}
	return int(ans % 1000000007)
}
