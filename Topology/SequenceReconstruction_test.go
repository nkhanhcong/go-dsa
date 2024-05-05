package topology

import (
	"fmt"
	"testing"
)

func TestSequenceReconstructionFalse(T *testing.T){
	nums := []int{1,2,3}
	sequences := [][]int{
		{1,2},
		{1,3},
	}

	res := SequenceReconstruction(nums, sequences)

	fmt.Println(res)

}

func TestSequenceReconstructionTrue(T *testing.T){
	nums := []int{1,2,3}
	sequences := [][]int{
		{1,2},
		{1,3},
		{1,2,3},
	}

	res := SequenceReconstruction(nums, sequences)

	fmt.Println(res)

}