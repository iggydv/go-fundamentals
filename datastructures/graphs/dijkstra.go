package graphs

func Dijkstra(graph map[int]map[int]int, start int) map[int]int {
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = -1
	}
	distances[start] = 0
	visited := make(map[int]bool)

	for len(visited) < len(graph) {
		node := minDistance(distances, visited)
		visited[node] = true

		for neighbor, weight := range graph[node] {
			if distances[neighbor] == -1 || distances[neighbor] > distances[node]+weight {
				distances[neighbor] = distances[node] + weight
			}
		}
	}
	return distances
}

func minDistance(distances map[int]int, visited map[int]bool) int {
	min := -1
	minNode := -1
	for node, distance := range distances {
		if distance != -1 && !visited[node] {
			if min == -1 || distance < min {
				min = distance
				minNode = node
			}
		}
	}
	return minNode
}

func generateAdjacencyList(edges [][]int) map[int]map[int]int {
	graph := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if _, ok := graph[u]; !ok {
			graph[u] = make(map[int]int)
		}
		if _, ok := graph[v]; !ok {
			graph[v] = make(map[int]int)
		}
		graph[u][v] = w
		graph[v][u] = w
	}
	return graph
}
