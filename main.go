package main

import (
	"algorithms/data-structures/unionfind/percolation"
	"fmt"
	"time"
)

// refactor so moving this into another area for UI consumption is possible
// use DI - may need refactor
// with UI - user inputs numGrids, height, width

type PercolationRunConfig struct {
	numberGrids int
	gridType    UnionFindDsType
	height      int
	width       int
}

func main() {
	/*
		numGrids := 1
		height := 500
		width := 500

		percolationRunConfigs := []PercolationRunConfig{
			PercolationRunConfig{
				numberGrids: numGrids,
				gridType:    unionfind.ArrayQuickFind,
				height:      height,
				width:       width,
			},
			PercolationRunConfig{
				numberGrids: numGrids,
				gridType:    unionfind.ArrayQuickUnion,
				height:      height,
				width:       width,
			},
		}

		for _, runConfig := range percolationRunConfigs {
			go func(config PercolationRunConfig) {
				s := make([]float64, config.numberGrids)
				for i := 0; i < config.numberGrids; i++ {
					go func(config PercolationRunConfig) {
						duration := execPercolationForGrid(config)
						fmt.Println(duration)
						s = append(s, duration)
					}(config)
				}
			}(runConfig)
		}

		fmt.Println("Completed")
	*/

	numGrids := 5
	height := 500
	width := 500

	percolation.StartSimulation(numGrids, height, width)

	/*
		c1 := make(chan []float64)
		c2 := make(chan []float64)

		var wg sync.WaitGroup
		wg.Add(percolationRunConfigs[0].numberGrids)
		wg.Add(percolationRunConfigs[1].numberGrids)

		go func(channel chan []float64) {
			t := time.Now()
			fmt.Printf("start time %v\n", t)
			config := percolationRunConfigs[0]
			s := make([]float64, 0, 100)
			for i := 0; i < config.numberGrids; i++ {
				go func() {
					d := execPercolationForGrid(config)
					s = append(s, d)
					wg.Done()
				}()
			}
			fmt.Printf("duration %v\n", time.Since(t))
			channel <- s
		}(c1)

		go func(channel chan []float64) {
			t := time.Now()
			fmt.Printf("start time %v\n", t)
			config := percolationRunConfigs[1]
			s := make([]float64, 0, 100)
			for i := 0; i < config.numberGrids; i++ {
				go func() {
					d := execPercolationForGrid(config)
					s = append(s, d)
					wg.Done()
				}()
			}
			fmt.Printf("duration %v\n", time.Since(t))
			channel <- s
		}(c2)

		wg.Wait()

		select {
		case qf := <-c1:
			fmt.Printf("Quick Find results: %v\n", qf)
			close(c1)
		case qu := <-c2:
			fmt.Printf("Quick Union results: %v\n", qu)
			close(c2)
		default:
			fmt.Println("Done")
		}
	*/
	// But of course you can have a result channel with buffer size len(params)
	// all goroutines send their answers and once your work is finished you collect from this result channel.
	/*
		c1 := make(chan []int)
		c2 := make(chan []int)

		go func(channel chan []int) {
			// there is a difference between time.Now vs time.Now()
			t := time.Now()
			fmt.Printf("start time %v\n", t)
			s := make([]int, 0, 100)
			for i := 0; i < 100; i++ {
				go func(i int) {
					s = append(s, i)
				}(i)
			}
			fmt.Printf("duration %v\n", time.Since(t))
			channel <- s
		}(c1)

		go func(channel chan []int) {
			t := time.Now()
			fmt.Printf("start time %v\n", t)
			s := make([]int, 0, 100)
			for i := 0; i < 100; i++ {
				go func(i int) {
					s = append(s, i)
				}(i)
			}
			fmt.Printf("duration %v\n", time.Since(t))
			channel <- s
		}(c2)

		fmt.Println(<-c1)
		fmt.Println(<-c2)

		close(c1)
		close(c2)
	*/
	/*
		numGrids := 10

		// no pauses at 100
		// 1s/0.5s 500
		// 25s/1.9s 800
		height := 500
		width := 500

		quickFindChannel := make(chan []float64)
		quickUnionChannel := make(chan []float64)

		go execPercolationForGridWithChannel(quickFindChannel, unionfind.ArrayQuickFind, numGrids, height, width)
		go execPercolationForGridWithChannel(quickUnionChannel, unionfind.ArrayQuickUnion, numGrids, height, width)

			s1 := <-quickFindChannel
			s2 := <-quickUnionChannel

			// print
			fmt.Println(s1)
			fmt.Println(s2)

				close(quickFindChannel)
				close(quickUnionChannel)
	*/
	/*
		select {
		case qf := <-quickFindChannel:
			fmt.Printf("Quick Find results: %v\n", qf)
		case qu := <-quickUnionChannel:
			fmt.Printf("Quick Union results: %v\n", qu)
		default:
			fmt.Println("Done")
		}
	*/
	// get the average run time
}

// essentially runs synchronously
func runSimulationUsingChannels() {
	numGrids := 1
	height := 500
	width := 500

	percolationRunConfigs := []PercolationRunConfig{
		PercolationRunConfig{
			numberGrids: numGrids,
			gridType:    unionfind.ArrayQuickFind,
			height:      height,
			width:       width,
		},
		PercolationRunConfig{
			numberGrids: numGrids,
			gridType:    unionfind.ArrayQuickUnion,
			height:      height,
			width:       width,
		},
	}

	fmt.Println(percolationRunConfigs[0].gridType)
	fmt.Println(percolationRunConfigs[1].gridType)

	for _, config := range percolationRunConfigs {
		fmt.Printf("Starting config : %v\n", config)
		s := make([]float64, 0, 100)
		for i := 0; i < config.numberGrids; i++ {
			c := make(chan float64)
			go func(channel chan float64) {
				fmt.Println(config)
				d := execPercolationForGrid(config)
				channel <- d
			}(c)
			duration := <-c
			fmt.Printf("Completed config : %v with duration : %v\n", config, duration)
			s = append(s, duration)
		}
	}

	fmt.Println("Complete")
}

func runSimulationUsingNestedChannels() {
	numGrids := 1
	height := 500
	width := 500

	percolationRunConfigs := []PercolationRunConfig{
		PercolationRunConfig{
			numberGrids: numGrids,
			gridType:    unionfind.ArrayQuickFind,
			height:      height,
			width:       width,
		},
		PercolationRunConfig{
			numberGrids: numGrids,
			gridType:    unionfind.ArrayQuickUnion,
			height:      height,
			width:       width,
		},
	}

	fmt.Println(percolationRunConfigs[0].gridType)
	fmt.Println(percolationRunConfigs[1].gridType)

	for _, config := range percolationRunConfigs {
		fmt.Printf("Starting config : %v\n", config)
		s := make([]float64, 0, 100)
		for i := 0; i < config.numberGrids; i++ {
			c := make(chan float64)
			go func(channel chan float64) {
				fmt.Println(config)
				d := execPercolationForGrid(config)
				channel <- d
			}(c)
			duration := <-c
			fmt.Printf("Completed config : %v with duration : %v\n", config, duration)
			s = append(s, duration)
		}
	}

	fmt.Println("Complete")
}

func execPercolationForGridWithChannel(c chan []float64, runConfig PercolationRunConfig) {
	s := make([]float64, runConfig.numberGrids)
	for i := 0; i < runConfig.numberGrids; i++ {
		go func(config PercolationRunConfig) {
			duration := execPercolationForGrid(config)
			s = append(s, duration)
		}(runConfig)
	}
	c <- s
}

func execPercolationForGrid(runConfig PercolationRunConfig) float64 {
	fmt.Printf("~~~ Creating grid : %v ~~~\n", runConfig.gridType)
	newGridFunc := unionfind.NewUnionFindGrid(runConfig.gridType)
	grid := newGridFunc(runConfig.height, runConfig.width)

	start := time.Now()
	fmt.Printf("Starting process at %v\n", start)
	iterations := grid.Process()
	fmt.Printf("Process took %d iterations to complete\n", iterations)

	duration := time.Since(start).Seconds()
	fmt.Printf("Completed in %v seconds\n", duration)
	fmt.Printf("~~~ Ending %v ~~~\n", runConfig.gridType)

	return duration
}
