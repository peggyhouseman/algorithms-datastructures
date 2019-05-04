package simulation

import (
	"fmt"
	"time"
)

type percolationResult struct {
	config  PercolationRunConfig
	results []float64
}

func StartPercolation(numGrids int, height int, width int) {

	runConfigs := NewPercolationRunConfigs(numGrids, height, width)

	var jobs []chan percolationResult

	for _, runConfig := range runConfigs {

		job := make(chan percolationResult)
		jobs = append(jobs, job)

		go func(configChannel chan percolationResult, config PercolationRunConfig) {
			var dChan []chan float64
			for i := 0; i < config.NumberGrids; i++ {
				c := make(chan float64)
				dChan = append(dChan, c)
				go execPercolationForGridWithChannel(c, config)
			}

			var s []float64
			for _, c := range dChan {
				d := <-c
				s = append(s, d)
				close(c)
			}

			result := percolationResult{
				config:  config,
				results: s,
			}

			configChannel <- result
		}(job, runConfig)
	}

	for _, job := range jobs {
		fmt.Println(<-job)
		close(job)
	}
}

func execPercolationForGridWithChannel(c chan<- float64, runConfig PercolationRunConfig) {
	fmt.Printf("~~~ Creating grid : %v ~~~\n", runConfig.GridType)
	newGridFunc := NewUnionFindGrid(runConfig.GridType)
	grid := newGridFunc(runConfig.Height, runConfig.Width)

	start := time.Now()
	fmt.Printf("Starting process at %v\n", start)
	iterations := grid.Process()
	fmt.Printf("Process took %d iterations to complete\n", iterations)

	duration := time.Since(start).Seconds()
	fmt.Printf("Completed in %v seconds\n", duration)
	fmt.Printf("~~~ Ending %v ~~~\n", runConfig.GridType)

	c <- duration
}
