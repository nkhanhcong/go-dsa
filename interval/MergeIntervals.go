package interval


func MergeInterval(intervals [][]int) [][]int{
	res := [][]int{}

	for _,interval := range intervals{

		if len(res) == 0 {
			res = append(res, interval)
		}else{
			lastInterval = res[len(res)-1]
			res = res[0:len(res)-1]
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
	return !IsLeftIntervals(a,b) || !IsRightIntervals(a,b)
}