package can_traverse_all_pairs

func canTraverseAllPairs(nums []int) bool {
	nodeMap := map[int][]int{}
	var gcd func(int, int) int
	gcd = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}
	//t1 := time.Now()
	nodeExistMap := map[int]bool{}
	newNums := []int{}
	for i := 0; i < len(nums); i++ {
		if !nodeExistMap[nums[i]] {
			newNums = append(newNums, nums[i])
			nodeExistMap[nums[i]] = true
		}
	}
	if len(newNums) == 1 {
		return true
	}
	for i := 0; i < len(newNums); i++ {
		for j := i + 1; j < len(newNums); j++ {
			if gcd(newNums[i], newNums[j]) > 1 {
				nodeMap[i] = append(nodeMap[i], j)
				nodeMap[j] = append(nodeMap[j], i)
			}
		}
	}
	//fmt.Println(time.Since(t1))
	if len(nodeMap[0]) == 0 {
		return false
	}
	stack := []int{0}
	visited := make([]bool, len(newNums))
	counter := 1
	//t1 = time.Now()
	for len(stack) != 0 {
		node := stack[0]
		visited[node] = true
		stack = stack[1:]
		for _, sonNode := range nodeMap[node] {
			if !visited[sonNode] {
				visited[sonNode] = true
				counter++
				stack = append(stack, sonNode)
			}
		}
	}
	//fmt.Println(time.Since(t1))
	return counter == len(newNums)
}
