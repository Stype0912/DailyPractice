package make_sum_divisible_by_p

func minSubarray(nums []int, p int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		} else {
			return j
		}
	}
	mod := func(i, j int) int {
		for i < 0 {
			i += j
		}
		return i % j
	}
	sum := 0
	preSumIndexMap := map[int]int{0: -1}
	preSum := []int{}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= p
		preSum = append(preSum, sum)
	}
	if sum%p == 0 {
		return 0
	}
	ans := -1
	for i := 0; i < len(nums); i++ {
		sumOfJ := mod(preSum[i]-sum+p, p)
		if j, ok := preSumIndexMap[sumOfJ]; !ok {

		} else {
			if i > j {
				if ans == -1 {
					ans = i - j
				} else {
					ans = min(ans, i-j)
				}
			}
		}
		preSumIndexMap[preSum[i]] = i
	}
	if ans >= len(nums) {
		return -1
	}
	return ans
}
