package score_of_parentheses

func scoreOfParentheses(s string) int {
	stack := []int{}
	stack = append(stack, 0)
	for _, item := range s {
		if item == '(' {
			stack = append(stack, 0)
		} else {
			element := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(1, 2*element)
		}
	}
	return stack[len(stack)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
