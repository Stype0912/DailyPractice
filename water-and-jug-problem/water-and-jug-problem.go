package water_and_jug_problem

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	stack := [][]int{}
	stack = append(stack, []int{0, 0})
	isSeen := map[[2]int]bool{}
	for {
		if len(stack) == 0 {
			break
		}
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		remainX, remainY := element[0], element[1]
		if remainX+remainY == targetCapacity || remainX == targetCapacity || remainY == targetCapacity {
			return true
		}
		if isSeen[[2]int{remainX, remainY}] == true {
			continue
		}
		isSeen[[2]int{remainX, remainY}] = true
		stack = append(stack, []int{jug1Capacity, remainY})
		stack = append(stack, []int{remainX, jug2Capacity})
		stack = append(stack, []int{0, remainY})
		stack = append(stack, []int{remainX, 0})
		if remainX+remainY < jug1Capacity {
			stack = append(stack, []int{remainX + remainY, 0})
		} else {
			stack = append(stack, []int{jug1Capacity, remainX + remainY - jug1Capacity})
		}
		if remainX+remainY < jug2Capacity {
			stack = append(stack, []int{0, remainX + remainY})
		} else {
			stack = append(stack, []int{remainX + remainY - jug2Capacity, 0})
		}
	}
	return false
}
