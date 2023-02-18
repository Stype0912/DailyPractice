package find_positive_integer_solution_for_a_given_equation

func findSolution(customFunction func(int, int) int, z int) [][]int {
	i, j := 1, 1000
	res := [][]int{}
	for {
		if i > 1000 || j <= 0 {
			break
		}
		if customFunction(i, j) == z {
			res = append(res, []int{i, j})
			i++
			j--
		} else if customFunction(i, j) > z {
			j--
		} else if customFunction(i, j) < z {
			i++
		}
	}
	return res
}
