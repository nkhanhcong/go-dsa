package dfs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCanTravelAllRoom(t *testing.T) {
	rooms := [][]int{[]int{1}, []int{2}, []int{3}, []int{}}

	res := CanVisitAllRooms(rooms)

	assert.Equal(t, true, res)

}


func TestCannotTravelAllRoom(t *testing.T) {
	rooms := [][]int{[]int{1,3}, []int{3,0,1}, []int{2}, []int{0}}

	res := CanVisitAllRooms(rooms)

	assert.Equal(t, false, res)

}
