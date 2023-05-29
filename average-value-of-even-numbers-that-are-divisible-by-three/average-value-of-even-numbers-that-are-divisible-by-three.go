package average_value_of_even_numbers_that_are_divisible_by_three

func averageValue(nums []int) int {
	cnt := 0
	sum := 0
	for _, num := range nums {
		if num%3 == 0 && num%2 == 0 {
			cnt += 1
			sum += num
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
