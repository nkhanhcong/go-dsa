package dfs

import "fmt"

func MinimumFuelCost(roads [][]int, seats int) int {

	graph := makeGraphCity(roads)

	res := 0

	var dfsCity func(node int, parrent int) int
	dfsCity = func(node int, parent int) int {

		totalPeople := 0

		for _, child := range graph[node] {

			if child != parent {

				people := dfsCity(child, node)
				totalPeople += people

				res += (totalPeople / seats)
				fmt.Println(totalPeople/seats, totalPeople, seats)
				if totalPeople%seats != 0 && parent != 1 {
					res += 1
				}
			}
		}

		return totalPeople + 1

	}

	dfsCity(0, -1)

	return res

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
