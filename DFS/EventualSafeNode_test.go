package dfs

import (
	"fmt"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	graph := [][]int{[]int{1, 2}, []int{2, 3}, []int{5}, []int{0}, []int{5}, []int{}, []int{}}
	res := EventualSafeNodes(graph)

	fmt.Println(res)
}
