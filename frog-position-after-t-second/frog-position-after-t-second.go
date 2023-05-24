package frog_position_after_t_second

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	nodeMap := map[int][]int{}
	for _, edge := range edges {
		nodeMap[edge[0]] = append(nodeMap[edge[0]], edge[1])
		nodeMap[edge[1]] = append(nodeMap[edge[1]], edge[0])
	}
	visited := map[int]bool{}
	var dfs func(int, int, float64)
	ans := 0.0
	dfs = func(node, jump int, prob float64) {
		if visited[node] || jump > t {
			return
		}
		if node == target && jump == t {
			ans = prob
			return
		}
		visited[node] = true
		sonNodes := nodeMap[node]
		sonNodeCount := 0
		for _, sonNode := range sonNodes {
			if !visited[sonNode] {
				sonNodeCount++
			}
		}
		if sonNodeCount == 0 && node == target {
			ans = prob
			return
		}
		for _, sonNode := range sonNodes {
			dfs(sonNode, jump+1, prob*1.0/float64(sonNodeCount))
		}
	}
	dfs(1, 0, 1)
	return ans
}
