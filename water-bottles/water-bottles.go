package main

func numWaterBottles(numBottles int, numExchange int) int {
	emptyBottles := 0
	ans := 0
	for {
		ans += numBottles
		tmp := numBottles
		numBottles = (tmp + emptyBottles) / numExchange
		emptyBottles = (tmp + emptyBottles) % numExchange
		if numBottles == 0 {
			return ans
		}
	}
}
