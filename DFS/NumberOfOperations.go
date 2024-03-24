package dfs


func MakeConnected(n int, connections [][]int) int {

	graph := makeGraph(connections, n)

	visitedMap := make(map[int]bool)
	dfsConnectionGraph(graph, 0, visitedMap)
	minLen := n - 1
	lenCable := len(visitedMap) - 1
	totalCable := len(connections)

	if totalCable < minLen {
		return -1
	}

	minCable := min(totalCable,minLen)


	return minCable - lenCable
}

func makeGraph(connections [][]int, n int) map[int][]int {
	graph := make(map[int][]int)

	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}
	return graph
}

func dfsConnectionGraph(graph map[int][]int, node int, visitedMap map[int]bool) {

	visitedMap[node] = true
	for _, nei := range graph[node] {
		if !visitedMap[nei] {
			dfsConnectionGraph(graph, nei, visitedMap)
		}
	}
}
