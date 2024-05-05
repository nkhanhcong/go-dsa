package topology

import "fmt"

func SequenceReconstruction(nums []int, sequences [][]int) bool {
	lengthNums := len(nums) + 1
	indegree := make([]int, lengthNums)
	graph := map[int][]int{}

	for _, s := range sequences {
		for i := 0; i < len(s)-1; {
			fmt.Println("i: ", i)
			u := s[i]
			v := s[i+1]
			indegree[v] += 1
			graph[u] = append(graph[u], v)
			i = i + 1

		}
	}

	queueSlice := []int{}

	for i, v := range indegree {
		if v == 0 && i != 0 {
			queueSlice = append(queueSlice, v)
		}
	}

	if len(queueSlice) != 1 {
		return false
	}

	fmt.Println("graph: ", graph)
	fmt.Println("indegree: ", indegree)
	for len(queueSlice) > 0 {
		q := queueSlice[0]
		queueSlice = queueSlice[1:]
		for _, s := range graph[q] {
			indegree[s] -= 1
			if indegree[s] == 0 {
				queueSlice = append(queueSlice, s)
			}
		}
		if len(queueSlice) > 1 {
			return false
		}
	}

	return true

}
