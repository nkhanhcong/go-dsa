package dfs

import (
	"fmt"
	"testing"
)


func TestDFSTravel( *testing.T){
	graph := CreateGraph()
	startNode := "0"

	visited:=DFSTravel(graph,startNode)

	fmt.Println(visited)

}