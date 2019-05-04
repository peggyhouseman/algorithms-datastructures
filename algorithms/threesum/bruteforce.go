package threesum

import (
	"cs/algo-ds/utils"
)

func FindThreeSumBruteForce(length int, sumValue int) []ThreeSumResult {
	arr := utils.NewIntArray(length)

	result := []ThreeSumResult{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == sumValue {
					result = append(result, ThreeSumResult{
						Value1: i,
						Value2: j,
						Value3: k,
					})
				}
			}
		}
	}

	return result
}
