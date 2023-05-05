package the_employee_that_worked_on_the_longest_task

func hardestWorker(n int, logs [][]int) int {
	employeeWorkTime := map[int]int{}
	employeeWorkTime[logs[0][0]] = logs[0][1]
	ansTime := logs[0][1]
	ansId := -1
	for i := 1; i < len(logs); i++ {
		employeeWorkTime[logs[i][0]] = logs[i][1] - logs[i-1][1]
		if employeeWorkTime[logs[i][0]] > ansTime {
			ansId = logs[i][0]
			ansTime = employeeWorkTime[logs[i][0]]
		} else if employeeWorkTime[logs[i][0]] == ansTime {
			if logs[i][0] < ansId {
				ansId = logs[i][0]
				ansTime = employeeWorkTime[logs[i][0]]
			}
		}
	}
	return ansId
}
