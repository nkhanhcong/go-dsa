package dfs

func MakeConnected(n int, connections [][]int) int {
	numConnection := len(connections)
	if numConnection < n-1 {
		return -1
	} 

	graph := makeGraph(connections)

	minConnected := 0
	visitedMap := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		if visitedMap[i] {
			continue
		}
		dfsConnectionGraph(graph, i, visitedMap)
		minConnected++
	}

	return minConnected - 1
}

func makeGraph(connections [][]int) map[int][]int {
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
