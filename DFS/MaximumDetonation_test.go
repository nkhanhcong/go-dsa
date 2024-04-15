package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumDenotation(t *testing.T) {
	bombs := [][]int{
		{2, 1, 3},
		{6, 1, 4},
	}

	assert.Equal(t, 2, MaximumDenotation(bombs))
}

func TestMaximumDenotationWithoutConnected(t *testing.T) {
	bombs := [][]int{
		{1, 1, 5},
		{10, 10, 5},
	}

	assert.Equal(t, 1, MaximumDenotation(bombs))

}

func TestMaximumDenotationWith5Connected(t *testing.T) {
	bombs := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 4, 2},
		{4, 5, 3},
		{5, 6, 4},
	}

	assert.Equal(t, 5, MaximumDenotation(bombs))
}
