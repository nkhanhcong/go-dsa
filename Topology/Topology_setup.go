package topology

func topologyDFS(graph [][]int) []int {
	visitedMap := map[int]bool{}
	stack := []int{}

	for i, _ := range graph {
		if !visitedMap[i] {
			stack = DFSForTopo(graph, visitedMap, i, stack)
		}
	}

	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack

}

func DFSForTopo(graph [][]int, visitedMap map[int]bool, node int, stack []int) []int {

	visitedMap[node] = true
	for _, neigh := range graph[node] {
		if !visitedMap[neigh] {
			stack = DFSForTopo(graph, visitedMap, neigh, stack)
		}
	}
	stack = append(stack, node)
	return stack
}
