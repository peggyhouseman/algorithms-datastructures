package main

import (
	"cs/algo-ds/unionfind/percolation"
)

// refactor so moving this into another area for UI consumption is possible
// use DI - may need refactor
// with UI - user inputs numGrids, height, width
func main() {
	numGrids := 5
	height := 500
	width := 500

	percolation.StartSimulation(numGrids, height, width)
}
