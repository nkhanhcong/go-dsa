package interval

import (
	"sort"
)

func MergeInterval(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
		} else {
			lastInterval := res[len(res)-1]
			res = res[0 : len(res)-1]
			if IsOverlapIntervals(lastInterval, interval) {
				newLastInterval := []int{
					min(lastInterval[0], interval[0]),
					max(lastInterval[1], interval[1]),
				}
				res = append(res, newLastInterval)
			} else {
				res = append(res, lastInterval)
				res = append(res, interval)
			}

		}
	}
	return res
}

func IsRightIntervals(a []int, b []int) bool {
	return a[1] < b[0]
}

func IsLeftIntervals(a []int, b []int) bool {
	return a[0] > b[1]
}

func IsOverlapIntervals(a []int, b []int) bool {
	return !IsLeftIntervals(a, b) && !IsRightIntervals(a, b)
}
