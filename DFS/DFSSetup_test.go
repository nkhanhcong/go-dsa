package dfs

import (
	"fmt"
	"testing"
)

func TestDFSTravel(*testing.T) {
	graph := CreateGraph()
	startNode := "0"

	visited := DFSTravel(graph, startNode)

	fmt.Println(visited)

}

func TestDFSRecursive(t *testing.T) {

	graph := CreateGraph()
	start := "0"

	visitedMap := make(map[string]bool)
	DFSTravelRecursive(graph, start, visitedMap)

	fmt.Println(visitedMap)
}

func TestDFSRecursive2(t *testing.T) {
	graph := CreateGraph()
	start := "0"

	visitedMap := make(map[string]bool)
	DFSTravelRecursive(graph, start, visitedMap)

	fmt.Println(visitedMap)
}
