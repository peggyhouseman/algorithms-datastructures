package simulation

import (
	"cs/algo-ds/algorithms/threesum"
	"fmt"
	"time"
)

func StartThreeSum(length int, sumValue int) {
	start := time.Now()
	fmt.Printf("~~~ Starting 3 sum process at %v ~~~\n", start)

	result := threesum.FindThreeSumBruteForce(length, sumValue)

	duration := time.Since(start).Seconds()
	fmt.Printf("~~~ Completed in %v seconds ~~~\n", duration)
	fmt.Printf("~~~ Result: %v ~~~\n", result)
}
