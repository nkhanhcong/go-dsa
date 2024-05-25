package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInterval(t *testing.T) {

	intervals := [][]int{
		{1, 3},
		{6, 9},
	}

	newInterval := []int{10, 11}

	res := InsertInterval(intervals, newInterval)
	expected := [][]int{
		{1, 3},
		{6, 9},
		{10, 11},
	}

	assert.Equal(t, expected, res)
}
