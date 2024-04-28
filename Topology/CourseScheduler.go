package topology

const (
	VISITED_CONSIDERED = 0
	VISITED_NOT_CYCLE  = -1
	VISITED_CYCLE      = 1
)

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraphCourses(numCourses, prerequisites)

	visitedMap := map[int]int{}
	for n, _ := range graph {
		if isCycle(graph, visitedMap, n) {
			return false
		}
	}

	return true
}

func isCycle(graph [][]int, visitedMap map[int]int, node int) bool {

	if visitedMap[node] == VISITED_CYCLE {
		return true
	}

	if visitedMap[node] == VISITED_NOT_CYCLE {
		return false
	}

	visitedMap[node] = VISITED_CYCLE

	for _, neigh := range graph[node] {
		if isCycle(graph, visitedMap, neigh) {
			return true
		}
	}
	visitedMap[node] = VISITED_NOT_CYCLE
	return false
}

func buildGraphCourses(numberCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numberCourses)

	for _, p := range prerequisites {
		n1, n2 := p[0], p[1]
		graph[n1] = append(graph[n1], n2)
	}
	return graph
}
