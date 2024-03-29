package dfs

import (
	"fmt"
)

func MinimumFuelCost(roads [][]int, seats int) int64 {
	graph := makeGraphCity(roads)
	visitedMap := make(map[int]bool)
	ans := int64(0)

	count, result := dfsCity(graph, 0, visitedMap, seats, ans)

	fmt.Println(count, ans)
	return int64(result)

}

func dfsCity(graph map[int][]int, node int, visited map[int]bool, seats int, ans int64) (int, int64) {

	_, exsited := visited[node]

	if exsited {
		return 0, int64(0)
	}

	visited[node] = true

	totalPeople := 0
	for _, neigh := range graph[node] {
		if visited[neigh] {
			continue
		}

		people, _ := dfsCity(graph, neigh, visited, seats,ans)
		ans += int64(people / seats)
		totalPeople += people

	}

	fmt.Println(totalPeople, node)

	return totalPeople + 1, ans

}

func makeGraphCity(roads [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, r := range roads {
		from, to := r[0], r[1]

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	return graph
}
