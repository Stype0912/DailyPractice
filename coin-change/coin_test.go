package coin_change

import "testing"

func TestCoin(t *testing.T) {
	t.Log(coinChange([]int{2, 5, 10, 1}, 27))
}
