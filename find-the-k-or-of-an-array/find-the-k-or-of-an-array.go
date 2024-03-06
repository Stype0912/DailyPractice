package find_the_k_or_of_an_array

func findKOr(nums []int, k int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				count++
			}
		}
		if count >= k {
			ans |= 1 << i
		}
	}
	return ans
}
