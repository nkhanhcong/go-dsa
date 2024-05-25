package interval

import "sort"

func FindMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	minArrows :=1 
	insectArrows := [][]int{points[0]}
	for _, p := range points[1:] {
		lastArrow := insectArrows[len(insectArrows)-1]

		if checkOverlap(p, lastArrow) {
			insectArrows[len(insectArrows)-1] = []int{
				max(p[0], lastArrow[0]),
				min(p[1], lastArrow[1]),
			}
		}else{
			minArrows +=1
			insectArrows = append(insectArrows, p)
		}

	}

	return minArrows
}

func checkOverlap(a []int, b []int) bool {
	isLeft := a[1] < b[0]
	isRight := a[0] > b[1]

	return !isLeft && !isRight
}
