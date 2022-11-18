package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	var stack []int
	profit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		num := prices[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			if stack[0]-num > profit {
				profit = stack[0] - num
			}
		}
		stack = append(stack, num)
	}
	return profit
}
