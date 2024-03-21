package dfs

import "fmt"

func IsBipartite(graph [][]int) bool {

	colorMap := make(map[int]string)
	return dfsNodeGraphBipartite(graph, colorMap, 0, "r")

}

func dfsNodeGraphBipartite(graph [][]int, colorMap map[int]string, node int, color string) bool {

	_, existed := colorMap[node]

	fmt.Println(graph, node, color, colorMap)
	if existed {
		return colorMap[node] != color
	}

	colorMap[node] = color
	for _,nei := range graph[node] {
		fmt.Println(node, nei,graph[node])
		if !dfsNodeGraphBipartite(graph, colorMap, nei, changeColor(color)) {
			return false
		}
	}
	return true

}

func changeColor(color string) string {
	if color == "r" {
		return "b"
	}

	return "r"
}
