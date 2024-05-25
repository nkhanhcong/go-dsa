package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinArrowShots2(t *testing.T) {

	// [[10,16],[2,8],[1,6],[7,12]]

	points := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}

	res := FindMinArrowShots(points)
	assert.Equal(t, res, 2)
}
