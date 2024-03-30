package dfs

import (
	"fmt"
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
}

func TestMinimumFuelCostWithNone(t *testing.T) {
	roads := [][]int{}

	seats := 2

	fmt.Println(MinimumFuelCost(roads, seats))
}
