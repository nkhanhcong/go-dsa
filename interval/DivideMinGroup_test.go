package interval

import (
	"fmt"
	"testing"
)

func TestMinGroups(t *testing.T) {

	// [[5,10],[6,8],[1,5],[2,3],[1,10]]
	intervals := [][]int{
		{5,10},
		{6,8},
		{1,5},
		{2,3},
		{1,10},
	}

	res := MinGroups(intervals)
	fmt.Println(res)
	
}