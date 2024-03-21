package dfs

import (
	"fmt"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	graph := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}

	fmt.Println(IsBipartite(graph))
}

func TestTrueBipartite(t *testing.T) {
	graph := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}

	fmt.Println(IsBipartite(graph))
}
