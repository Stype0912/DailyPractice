package validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for i := 0; i < len(popped); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != popped[i] {
			flag := false
			for k := j; k < len(pushed); k++ {
				if pushed[k] == popped[i] {
					j = k + 1
					flag = true
					break
				}
				stack = append(stack, pushed[k])
			}
			if !flag {
				return false
			}
		}
		if popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
