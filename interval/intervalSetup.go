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

func IsContained(a []int, b []int )bool{
	return a[0] <= b[0] && a[1]>= b[1]
}

func Overlap(a []int, b []int) bool{
	return !IsLeft(a,b) && !IsRight(a,b)
}

func NotOverlap(a []int, b[]int) bool{
	return IsLeft(a,b) || IsRight(a,b)
}

func Intersect(a []int, b []int) []int{

	if NotOverlap(a,b){
		return []int{} 
	}

	return []int{max(a[0],b[0]), min(a[1],b[1])}
}

func Merge( a[]int, b []int) []int{


	return []int{}
} 