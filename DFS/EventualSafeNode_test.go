package dfs

import (
	"fmt"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	res := EventualSafeNodes(graph)

	fmt.Println(res)

	graph1 := [][]int{
		{1, 2, 3, 4},
		{1, 2},
		{3, 4},
		{0, 4},
		{},
	}

	res1 := EventualSafeNodes(graph1)
	fmt.Println(res1)

}
