package dfs

import (
	"fmt"
	"math"
)

func MaximumDenotation(bombs [][]int) int {

	graph := buildGraphBomb(bombs)
	fmt.Println(graph)

	return 0
}

func buildGraphBomb(bombs [][]int) map[int][]int {

	graph := make(map[int][]int)
	for i, b := range bombs {

		for j := i + 1; j < len(bombs); j++ {

			if isConnectedBomb(b, bombs[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}

		}
	}
	return graph
}

func isConnectedBomb(bombA []int, bombB []int) bool {

	xA, yA, rA := bombA[0], bombA[1], bombA[2]

	xB, yB, rB := bombB[0], bombB[1], bombB[2]

	distanceAB := math.Sqrt(math.Pow(float64(xA)-float64(xB), 2) + math.Pow(float64(yA)-float64(yB), 2))

	return distanceAB < float64(rA+rB)
}
