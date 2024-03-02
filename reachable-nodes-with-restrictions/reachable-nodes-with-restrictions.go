package reachable_nodes_with_restrictions

func reachableNodes(n int, edges [][]int, restricted []int) int {
	route := make([][]int, n)
	restrictedMap := make([]bool, n)
	visited := make([]bool, n)
	for i := 0; i < len(edges); i++ {
		node1, node2 := edges[i][0], edges[i][1]
		route[node1] = append(route[node1], node2)
		route[node2] = append(route[node2], node1)
	}
	for _, restrictedNode := range restricted {
		restrictedMap[restrictedNode] = true
	}
	var dfs func(node int)
	var res int = 0
	dfs = func(node int) {
		visited[node] = true
		res++
		for _, son := range route[node] {
			if !visited[son] && !restrictedMap[son] {
				dfs(son)
			}
		}
	}
	dfs(0)
	return res
}
