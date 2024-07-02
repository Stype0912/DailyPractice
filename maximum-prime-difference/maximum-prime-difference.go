package maximum_prime_difference

func maximumPrimeDifference(nums []int) int {
	primeNums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	primeSet := map[int]bool{}
	for _, num := range primeNums {
		primeSet[num] = true
	}
	ans := 0
	first := -1
	for i, num := range nums {
		if primeSet[num] {
			if first < 0 {
				first = i
			} else {
				ans = i - first
			}
		}
	}
	return ans
}
