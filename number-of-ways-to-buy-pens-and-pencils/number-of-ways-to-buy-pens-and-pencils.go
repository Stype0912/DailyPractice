package number_of_ways_to_buy_pens_and_pencils

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	maxPen := total / cost1
	ans := 0
	for i := 0; i <= maxPen; i++ {
		left := total - i*cost1
		ans += 1 + (left / cost2)
	}
	return int64(ans)
}
