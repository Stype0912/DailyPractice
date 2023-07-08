package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
