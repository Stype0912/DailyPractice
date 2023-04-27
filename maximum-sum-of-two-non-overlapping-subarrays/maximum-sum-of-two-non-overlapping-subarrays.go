package maximum_sum_of_two_non_overlapping_subarrays

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	dpFirstLeft := map[int]int{}
	dpFirstRight := map[int]int{}
	dpSecondLeft := map[int]int{}
	dpSecondRight := map[int]int{}
	tmp1, tmp2 := 0, 0
	for i := 0; i < len(nums); i++ {
		tmp1 += nums[i]
		tmp2 += nums[i]
		if i < firstLen-1 {
		} else if i == firstLen-1 {
			dpFirstLeft[i] = tmp1
		} else {
			tmp1 -= nums[i-firstLen]
			dpFirstLeft[i] = max(dpFirstLeft[i-1], tmp1)
		}
		if i < secondLen-1 {
		} else if i == secondLen-1 {
			dpSecondLeft[i] = tmp2
		} else {
			tmp2 -= nums[i-secondLen]
			dpSecondLeft[i] = max(dpSecondLeft[i-1], tmp2)
		}
	}
	tmp1, tmp2 = 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		tmp1 += nums[i]
		tmp2 += nums[i]
		if len(nums)-i < firstLen {
		} else if len(nums)-i == firstLen {
			dpFirstRight[i] = tmp1
		} else {
			tmp1 -= nums[i+firstLen]
			dpFirstRight[i] = max(dpFirstRight[i+1], tmp1)
		}
		if len(nums)-i < secondLen {
		} else if len(nums)-i == secondLen {
			dpSecondRight[i] = tmp2
		} else {
			tmp2 -= nums[i+secondLen]
			dpSecondRight[i] = max(dpSecondRight[i+1], tmp2)
		}
	}
	ans := 0
	for i := 0; i <= len(nums); i++ {
		ans = max(ans, max(dpFirstLeft[i-1]+dpSecondRight[i], dpSecondLeft[i-1]+dpFirstRight[i]))
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
