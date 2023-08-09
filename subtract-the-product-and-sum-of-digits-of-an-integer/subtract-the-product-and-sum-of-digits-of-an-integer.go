package subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	mul := 1
	sum := 0
	for n != 0 {
		tmp := n % 10
		mul *= tmp
		sum += tmp
		n /= 10
	}
	return mul - sum
}
