package topology

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinHeightTree(t *testing.T){

	n := 4
	edges := [][]int{
		{1,0},
		{1,2},
		{1,3},
	}
	rootList := FindMinHeightTree(n, edges)
	assert.Equal(t, rootList, []int{1})
}

func TestFindMinHeightTreeWith6(t *testing.T){

	n := 6 
	edges := [][]int{
		{3,0},
		{3,1},
		{3,2},
		{3,4},
		{5,4},
	}
	rootList := FindMinHeightTree(n, edges)
	assert.Equal(t,rootList,[]int{3,4})
}