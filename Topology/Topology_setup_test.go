package topology

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologySetup(t *testing.T) {
	fmt.Println("Topology Setup Test")
	adjancyList := [][]int{
		{2, 4},
		{3},
		{1},
		{},
		{1, 2, 5},
		{3},
	}

	stack := topologyDFS(adjancyList)

	expectedOutput := []int{0, 4, 5, 2, 1, 3}

	assert.Equal(t, expectedOutput, stack)
}

func TestCyclicTopo(t *testing.T) {
	adjancyList := [][]int{
		{1},
		{0},
	}

	stack := topologyDFS(adjancyList)
	fmt.Println(stack)
}
