package make_costs_of_paths_equal_in_a_binary_tree

func minIncrements(n int, cost []int) int {
	ans := 0
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := n - 2; i > 0; i -= 2 {
		ans += abs(cost[i] - cost[i+1])
		cost[i/2] = cost[i/2] + max(cost[i], cost[i+1])
	}
	return ans
}
