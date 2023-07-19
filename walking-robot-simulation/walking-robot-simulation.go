package walking_robot_simulation

func robotSim(commands []int, obstacles [][]int) int {
	direction := 0
	now := []int{0, 0}
	movement := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	obstacleMap := map[[2]int]bool{}
	for i := 0; i < len(obstacles); i++ {
		obstacleMap[[2]int{obstacles[i][0], obstacles[i][1]}] = true
	}
	ans := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for _, command := range commands {
		if command == -1 {
			direction = (direction + 1) % 4
		} else if command == -2 {
			direction = (direction + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				if obstacleMap[[2]int{now[0] + movement[direction][0], now[1] + movement[direction][1]}] {
					break
				}
				now = []int{now[0] + movement[direction][0], now[1] + movement[direction][1]}
				ans = max(ans, now[0]*now[0]+now[1]*now[1])
			}
		}
	}
	return ans
}
