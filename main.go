package main

import "cs/algo-ds/simulation"

type test struct {
	number int
}

// refactor so moving this into another area for UI consumption is possible
// use DI - may need refactor
// with UI - user inputs numGrids, height, width
func main() {

	numGrids := 5
	height := 500
	width := 500

	simulation.StartPercolation(numGrids, height, width)

	length := 200
	sum := 40

	simulation.StartThreeSum(length, sum)

}
