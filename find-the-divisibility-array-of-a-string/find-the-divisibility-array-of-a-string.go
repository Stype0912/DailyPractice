package find_the_divisibility_array_of_a_string

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	num := 0
	for i := 0; i < len(word); i++ {
		num = num*10%m + int(word[i]-'0')
		if num%m == 0 {
			ans[i] = 1
		}
	}
	return ans
}
