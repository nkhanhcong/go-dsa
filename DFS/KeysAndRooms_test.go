package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTravelAllRoom(t *testing.T) {
	rooms := [][]int{{1}, {2}, {3}, {}}

	res := CanVisitAllRooms(rooms)

	assert.Equal(t, true, res)

}

func TestCannotTravelAllRoom(t *testing.T) {
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}

	res := CanVisitAllRooms(rooms)

	assert.Equal(t, false, res)

}

