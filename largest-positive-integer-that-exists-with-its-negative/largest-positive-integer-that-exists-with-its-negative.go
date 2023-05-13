package largest_positive_integer_that_exists_with_its_negative

func findMaxK(nums []int) int {
	numMap := map[int]int{}
	ans := -1
	for _, num := range nums {
		numMap[num]++
		if numMap[-1*num] != 0 {
			if num > 0 && num > ans {
				ans = num
			} else if num < 0 && -1*num > ans {
				ans = -1 * num
			}

		}
	}
	return ans
}
