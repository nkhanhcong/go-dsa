package topology

func TopologyBFS(graph [][]int, v int)[]int{
	indegree := make([]int, v)
	for k,_ := range graph{
		for _,vj := range graph[k]{
			indegree[vj]++
		}
	}

	queue := []int{}
	for _,v := range indegree{
		if v == 0{
			queue = append(queue, v)
		}
	}

	res := []int{}
	for len(queue)>0 {
		currentE := queue[0]
		res= append(res, currentE)
		queue = queue[1:]
		for _,v := range graph[currentE]{
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(res) != v{
		return []int{} 
	}

	return res 
}