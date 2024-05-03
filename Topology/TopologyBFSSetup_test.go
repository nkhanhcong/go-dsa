package topology

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	stack := TopologyBFS(adjancyList, 6)

	expectedOutput := []int{0, 4, 2, 5, 1, 3}
	assert.Equal(t, expectedOutput,stack)
}