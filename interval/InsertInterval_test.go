package interval

import (
	"fmt"
	"testing"
)

func TestInsertInterval(t *testing.T){

	intervals := [][]int{
		{1,3},
		{6,9},
	}

	newInterval := []int{10,11}


	res := InsertInterval(intervals,newInterval)

	fmt.Println(res)
}