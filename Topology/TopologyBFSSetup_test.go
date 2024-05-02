package topology

import (
	"fmt"
	"testing"
)


func TestTopologyBFSSetup(t *testing.T) {
	fmt.Println("Topology Setup Test")
	adjancyList := [][]int{
		{2, 4},
		{3},
		{1},
		{},
		{1, 2, 5},
		{3},
	}

	stack := TopologyBFS(adjancyList, 5)
	fmt.Println(stack)

	// expectedOutput := []int{0, 4, 5, 2, 1, 3}
}