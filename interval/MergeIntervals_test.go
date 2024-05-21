package interval

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	res := MergeInterval(intervals)

	fmt.Println(res)
}

func TestMergeWithOneArrayResult(t *testing.T) {
	intervals := [][]int{
		{2,3},
		{4,5},
		{6,7},
		{8,9},
		{1,10},
	}

	res := MergeInterval(intervals)

	fmt.Println(res)
	
}


func TestMergeWithBeginningAndLast(t *testing.T) {
	intervals := [][]int{
		{2,3},
		{5,5},
		{2,2},
		{3,4},
		{3,4},
	}
	res := MergeInterval(intervals)

	fmt.Println(res)
	
}