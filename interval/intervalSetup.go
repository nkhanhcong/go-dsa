package interval

// func NotOverlap(a []int, b []int)[]int{
	
// }

func IsLeft(a []int, b []int) bool{
	return a[1] < b[0]
}

func IsRight(a []int, b []int)bool{
	return a[0] > b[1]
}

func Contains(a []int, b []int )bool{
	return a[0] >= b[0] && a[1] <= b[1]
}