package best_time_to_buy_and_sell_stock_ii

import "testing"

func TestMaxProfit(t *testing.T) {
	profit := maxProfit([]int{1, 5, 3, 9999, 10000})
	t.Log(profit)
}
