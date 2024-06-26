package course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := map[int][]int{}
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}
	var dfs func(int)
	var stack []int
	var isValid bool = true
	status := map[int]int{}
	dfs = func(i int) {
		status[i] = 1
		for _, preCourse := range edges[i] {
			if status[preCourse] == 0 {
				dfs(preCourse)
			} else if status[preCourse] == 1 {
				isValid = false
			}
		}
		status[i] = 2
		stack = append(stack, i)
	}
	for i := 0; i < numCourses; i++ {
		if status[i] == 0 {
			dfs(i)
		}
	}
	if !isValid {
		return []int{}
	}
	ans := []int{}
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}
