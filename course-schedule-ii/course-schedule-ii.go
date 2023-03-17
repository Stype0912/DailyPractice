package course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	stack := []int{}
	status := map[int]int{}
	var dfs func(int)
	var valid bool = true
	dfs = func(i int) {
		if status[i] != 2 {
			status[i] = 1
		}
		for _, preRequisite := range prerequisites {
			if preRequisite[1] == i {
				if status[preRequisite[0]] == 0 {
					dfs(preRequisite[0])
				} else if status[preRequisite[0]] == 1 {
					valid = false
				}
			}
		}
		if status[i] <= 1 {
			status[i] = 2
			stack = append(stack, i)
		}
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	if !valid {
		return []int{}
	}
	ans := []int{}
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}
