package longest_valid_parentheses

func longestValidParentheses(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	maxLen := 0
	for i, letter := range s {
		if letter == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
