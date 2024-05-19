package interval

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOverlap(t *testing.T){

	a := []int{4,6}
	b := []int{5,8}
	res := Overlap(a,b)

	assert.Equal(t, res, true)
}


