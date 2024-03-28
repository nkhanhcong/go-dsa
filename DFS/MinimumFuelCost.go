package dfs


func MinimumFuelCost(roads [][]int, seats int) int64{
	graph := makeGraphCity(roads)
	visitedMap := make(map[int]bool)

	count:= dfsCity(graph,0, visitedMap,0)

	return int64(count)

}

func dfsCity(graph map[int][]int, node int, visited map[int]bool, count int)int{

	visited[node] = true

	for _, neigh := range graph[node]{
		if !visited[neigh]{
			dfsCity(graph,neigh,visited,count+1)
		}
	}

	return 0

}

func makeGraphCity(roads [][]int) map[int][]int{
	graph := make(map[int][]int)

	for _, r := range roads{
		from, to := r[0], r[1]

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	return graph
}