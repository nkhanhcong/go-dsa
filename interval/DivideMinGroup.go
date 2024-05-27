package interval

import "sort"

func MinGroups(intervals [][]int) int {
	resQueueLast := []int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		isChanged := false
		for i, res := range resQueueLast {
			start := interval[0]
			if res < start {
				resQueueLast[i] = interval[1]
				isChanged = true
				break
			}
		}
		if !isChanged {
			resQueueLast = append(resQueueLast, interval[1])
		}
	}

	return len(resQueueLast)
}
