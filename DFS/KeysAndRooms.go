package dfs

func CanVisitAllRooms(rooms [][]int) bool {
	stackLockedRooms := []int{0}
	visitedRooms := map[int]bool{0: true}

	for len(stackLockedRooms) > 0 {

		room := stackLockedRooms[len(stackLockedRooms)-1]
		stackLockedRooms = stackLockedRooms[:len(stackLockedRooms)-1]

		for _, nextRoom := range rooms[room] {
			if !visitedRooms[nextRoom] {
				visitedRooms[nextRoom] = true
				stackLockedRooms = append(stackLockedRooms, nextRoom)
			}

		}

	}

	return len(rooms) == len(visitedRooms)
}
