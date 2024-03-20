package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCircleNum(t *testing.T) {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	res := FindCircleNum(isConnected)

	assert.Equal(t, 2, res)
}
