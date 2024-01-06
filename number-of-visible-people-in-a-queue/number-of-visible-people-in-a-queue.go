package number_of_visible_people_in_a_queue

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= heights[i] {
			stack = stack[:len(stack)-1]
			res[i]++
		}
		if len(stack) > 0 {
			res[i]++
		}
		stack = append(stack, heights[i])
	}
	return res
}
