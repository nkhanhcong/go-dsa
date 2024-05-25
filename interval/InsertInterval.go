package interval

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	merged := false

	for _, interval := range intervals {
		if newInterval[1] < interval[0] {
			if !merged {
				res = append(res, newInterval)
				merged = true
			}
			res = append(res, interval)
		} else if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	if !merged {
		res = append(res, newInterval)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
