package dfs

import "fmt"

func CreateGraph() map[string][]string {

	graph := make(map[string][]string)
	graph["0"] = []string{"1", "2", "3"}
	graph["1"] = []string{"0", "2"}
	graph["2"] = []string{"0", "1", "4"}
	graph["3"] = []string{"0"}
	graph["4"] = []string{"2"}

	return graph
}

func DFSTravel(graph map[string][]string, start string) map[string]bool {

	stackSlice := []string{start}
	visitedMap := map[string]bool{start: true}

	for len(stackSlice) > 0 {
		node := stackSlice[len(stackSlice)-1]
		stackSlice = stackSlice[:len(stackSlice)-1]

		for _, neighbor := range graph[node] {
			if !visitedMap[neighbor] {
				visitedMap[neighbor] = true
				stackSlice = append(stackSlice, neighbor)
			}
		}
	}

	return visitedMap
}

func DFSTravelRecursive(graph map[string][]string, start string, visited map[string]bool) {
	visited[start] = true

	fmt.Println(start)
	for _, nei := range graph[start] {
		if !visited[nei] {
			DFSTravelRecursive(graph, nei, visited)
		}
	}

}
