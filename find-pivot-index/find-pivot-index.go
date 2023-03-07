package find_pivot_index

func pivotIndex(nums []int) int {
	sumMap := map[int]int{}
	sum := 0
	for index, num := range nums {
		sum += num
		sumMap[index] = sum
	}
	for index, num := range nums {
		if sumMap[index]-num == sum-sumMap[index] {
			return index
		}
	}
	return -1
}
