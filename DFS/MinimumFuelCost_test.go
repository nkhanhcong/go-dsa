package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumFuelCost(t *testing.T) {
	roads := [][]int{
		{3, 1},
		{3, 2},
		{1, 0},
		{0, 4},
		{0, 5},
		{4, 6},
	}

	seats := 2

	ans := MinimumFuelCost(roads, seats)

	assert.Equal(t, 7, ans)

	ans2 := MinimumFuelCost2(roads, seats)
	assert.Equal(t, int64(7), ans2)
}

func TestMinimumFuelCostWithNone(t *testing.T) {
	roads := [][]int{}

	seats := 2

	assert.Equal(t, 0, MinimumFuelCost(roads, seats))
	assert.Equal(t, int64(0), MinimumFuelCost2(roads, seats))
}

func TestTravelOnlyOneOnEachNode(t *testing.T) {
	roads := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
	}

	seats := 5

	assert.Equal(t, 3, MinimumFuelCost(roads, seats))
	assert.Equal(t, int64(3), MinimumFuelCost2(roads, seats))
}
