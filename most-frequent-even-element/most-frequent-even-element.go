package most_frequent_even_element

func mostFrequentEven(nums []int) int {
	numCount := map[int]int{}
	for _, num := range nums {
		if num%2 == 0 {
			numCount[num]++
		}
	}
	ans := -1
	count := 0
	for key, value := range numCount {
		if value > count {
			count = value
			ans = key
		} else if value == count {
			if key < ans {
				ans = key
			}
		}
	}
	return ans
}
