package dijkstra

const INF = 0x3f3f3f3f

// 迪杰斯特拉函数
func dijkstra(graph [][]int, start int) []int {
	n := len(graph)         // 图中顶点个数
	visit := make([]int, n) // 标记已经作为中间结点完成访问的顶点
	dist := make([]int, n)  // 存储从前点到其他顶点的最短路径

	for i := 0; i < n; i++ {
		dist[i] = graph[start][i] // 初始化遍历起点
	}
	visit[start] = 1 // 标记初始顶点

	var minDist, midNode int

	// 更新其他顶点最短路径，循环n次
	for i := 0; i < n; i++ {
		minDist = INF // 存储从起点到其他未被访问的结点中的最短路径
		midNode = 0   // 存储最短路径的结点编号

		// 遍历n个顶点，寻找未被访问且距离为起始位置到该点距离最小的顶点
		for j := 0; j < n; j++ {
			if visit[j] == 0 && minDist > dist[j] {
				minDist = dist[j] // 更新未被访问结点的最短路径
				midNode = j       // 更新顶点编号
			}
		}
		// 以midNode为中间结点，再循环遍历其他节点更新最短路径
		for j := 0; j < n; j++ {
			// 若该节点未被访问且找到更短路径即更新最短路径
			if visit[j] == 0 && dist[j] > dist[midNode]+graph[midNode][j] {
				dist[j] = dist[midNode] + graph[midNode][j]
			}
		}
		visit[midNode] = 1 // 标记已访问
	}
	return dist
}
