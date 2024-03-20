package dfs

import "fmt"

func EventualSafeNodes(graph [][]int) []int {
	visitedMap := make(map[int]bool)
	stateMap := make(map[int]int)

	for nNode := range graph {
		dfsNode(graph, nNode, visitedMap, stateMap)
		
	}

	fmt.Println("this is stateMap", stateMap)
	res := make([]int, 0)

	for k, v := range stateMap {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

// None is node not travel yet
// 1 is node is in stack
// 2 is node is safe node

func dfsNode(graph [][]int, start int, visitedMap map[int]bool, stateMap map[int]int) {
	stack := []int{start}

	visitedMap[start] = true
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(graph[node]) == 0 {
			stateMap[node] = 2 
		} else {
			stateMap[node] = 1
		}

		for _, nextNode := range graph[node] {
			if !visitedMap[nextNode] {
				visitedMap[nextNode] = true
				stateMap[nextNode] = 1
				stack = append(stack, nextNode)
			} else{
				if len(graph[nextNode]) ==0 {
					stateMap[node] = 2
				}else{
					stateMap[node] = 1
				}
			}
		}

	}

}
