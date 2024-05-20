package interval

import "fmt"


func MergeInterval(intervals [][]int) [][]int{
	res := [][]int{}

	for _,interval := range intervals{

		if len(res) == 0 {
			res = append(res, interval)
		}else{
			lastInterval := res[len(res)-1]
			res = res[0:len(res)-1]
			if IsOverlapIntervals(interval, lastInterval){
				left := min(interval[0], lastInterval[0])
				right := max(interval[1],lastInterval[1])
				res = append(res, []int{left,right})
				for len(res)> 1 {
					lastInterval := res[len(res) -1]
					nearLastInterval := res[len(res)-2]
					res = res[0:len(res)-2]
					if IsOverlapIntervals(lastInterval,nearLastInterval){
						left := min(lastInterval[0], nearLastInterval[0])
						right := max(lastInterval[1], nearLastInterval[1])
						res = append(res, []int{left,right})
					}else{
						break
					}
				}
			}else{
				res = append(res, lastInterval)
				res = append(res, interval)
			}
			
		}

	}

	return res
}

func IsRightIntervals(a []int, b []int) bool{
	return a[1] <= b[0]
}

func IsLeftIntervals(a []int, b []int) bool{
	return a[0] >= b[1]
}

func IsOverlapIntervals(a []int, b[]int)bool{
	return !IsLeftIntervals(a,b) && !IsRightIntervals(a,b)
}