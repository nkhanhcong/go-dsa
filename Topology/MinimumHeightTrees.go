package topology

func FindMinHeightTree(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjGraph := map[int][]int{}

	for _, e := range edges {
		u, v := e[0], e[1]
		adjGraph[u] = append(adjGraph[u], v)
		adjGraph[v] = append(adjGraph[v], u)
	}

	leavesQueue := []int{}
	indegree := make([]int, n)

	for k, v := range adjGraph {
		indegree[k] = len(v)
		if indegree[k] == 1 {
			leavesQueue = append(leavesQueue, k)
		}
	}

	for len(leavesQueue) > 0 {

		if n <= 2 {
			return leavesQueue
		}
		for _,leavesNode := range leavesQueue{
			leavesQueue = leavesQueue[1:]
			n -= 1
			for _, nei := range adjGraph[leavesNode] {
				indegree[nei] -= 1
				if indegree[nei] == 1 {
					leavesQueue = append(leavesQueue, nei)
				}
			}
		}

	}

	return []int{0}

}
