package maximum_odd_binary_number

func maximumOddBinaryNumber(s string) string {
	ans := ""
	nums := [2]int{}
	for _, num := range s {
		nums[num-'0']++
	}
	for i := 0; i < nums[1]-1; i++ {
		ans += "1"
	}
	for i := 0; i < nums[0]; i++ {
		ans += "0"
	}
	return ans + "1"
}
