package dfs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeConnected(t *testing.T) {

	n := 4
	connections := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}

	res := MakeConnected(n, connections)

	assert.Equal(t, 1, res)

}

func TestCannotConnected(t *testing.T) {
	n := 6
	connections := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
	}

	res := MakeConnected(n, connections)
	assert.Equal(t, -1, res)

}

func TestNeed2Connected(t *testing.T) {
	n := 6
	connections := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
	}

	res := MakeConnected(n, connections)

	assert.Equal(t, 2, res)
}

func TestConnection12(t *testing.T) {
	n := 12
	connections := [][]int{
		{1, 5},
		{1, 7},
		{1, 2},
		{1, 4},
		{3, 7},
		{4, 7},
		{3, 5},
		{0, 6},
		{0, 1},
		{0, 4},
		{2, 6},
		{0, 3},
		{0, 2},
	}

	res := MakeConnected(n, connections)
	fmt.Println(res)
}
