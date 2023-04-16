package find_the_maximum_divisibility_score

func maxDivScore(nums []int, divisors []int) int {
	ans := divisors[0]
	max := 0
	for _, divisor := range divisors {
		count := 0
		for _, num := range nums {
			if num%divisor == 0 {
				count++
			}
		}
		if count > max {
			max = count
			ans = divisor
		} else if count == max {
			if divisor < ans {
				ans = divisor
			}
		}
	}
	return ans
}
