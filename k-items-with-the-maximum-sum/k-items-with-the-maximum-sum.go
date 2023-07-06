package k_items_with_the_maximum_sum

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	ans := 0
	for k > 0 {
		k--
		if numOnes > 0 {
			ans++
			numOnes--
			continue
		}
		if numZeros > 0 {
			numZeros--
			continue
		}
		if numNegOnes > 0 {
			numNegOnes--
			ans--
			continue
		}
	}
	return ans
}
