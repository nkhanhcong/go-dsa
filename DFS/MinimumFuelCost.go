package dfs

import (
	"math"
)

func MinimumFuelCost(roads [][]int, seats int) int {

	graph := makeGraphCity(roads)
	var dfsCity func(node int, parrent int) (int, int)
	dfsCity = func(node int, parent int) (int, int) {

		totalPeople := 1
		localCost := 0

		for _, child := range graph[node] {

			if child != parent {

				people, cost := dfsCity(child, node)
				totalPeople += people

				localCost += cost
			}
		}
		trips := totalPeople / seats

		if totalPeople%seats != 0 {
			trips++
		}

		if node != 0 {
			localCost += trips

		}

		return totalPeople, localCost

	}

	_, res := dfsCity(0, -1)

	return res

}

func MinimumFuelCost2(roads [][]int, seats int) int64 {
	graph := makeGraphCity(roads)
	res := int64(0)

	var dfsCity func(node int, parent int) int
	dfsCity = func(node int, parent int) int {
		totalPeople := 1
		for _, child := range graph[node] {
			if child != parent {
				people := dfsCity(child, node)
				totalPeople += people
			}
		}

		if node != 0 {
			res += int64(math.Ceil(float64(totalPeople) / float64(seats)))
		}
		return totalPeople

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
