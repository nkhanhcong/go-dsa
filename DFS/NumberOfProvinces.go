package dfs

func FindCircleNum(isConnected [][]int) int {

	visited := make(map[int]bool)
	count := 0

	for cNumber, _ := range isConnected {

		if !visited[cNumber] {
			count += 1
			dfs(visited, cNumber, isConnected)

		}
	}

	return count
}

func dfs(visited map[int]bool, start int, isConnected [][]int) {
	stack := []int{start}

	for len(stack) > 0 {
		city := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for cNumber, v := range isConnected[city] {
			if !visited[cNumber] && v == 1 {
				visited[cNumber] = true
				stack = append(stack, cNumber)
			}
		}
	}
}
