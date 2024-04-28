package topology

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCourseScheduleCanFinish(t *testing.T) {
	prerequisites := [][]int{
		{1, 0},
	}

	res := CanFinish(2, prerequisites)
	assert.Equal(t, true, res)

}

func TestCourseScheduleCannotFinish(t *testing.T) {
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}

	res := CanFinish(2, prerequisites)

	assert.Equal(t, false, res)
}
