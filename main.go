package main

import (
	"math"
)


func Solution(A []int) int {
	allcombs := possibleCombs(A)
	lenLargetSubsetANDisPositive := 0
	for i := len(allcombs) - 1; i >= 0; i-- {
		if qualify(allcombs[i]) {
			lenLargetSubsetANDisPositive = len(allcombs[i])
			break
		}
	}
	return lenLargetSubsetANDisPositive
}

func possibleCombs(arr []int) [][]int {
	var combs [][]int
	total := int(math.Pow(2, float64(len(arr))))

	for i := 0; i < total; i++ {
		var comb []int
		for j := 0; j < len(arr); j++ {
			if i&(1<<j) != 0 {
				comb = append(comb, arr[j])
			}
		}
		if len(comb) > 0 {
			combs = append(combs, comb)
		}
	}
	return combs
}

func qualify(binarysubset []int) bool {
	if len(binarysubset) == 1 {
		return true
	}
	cumulativeAND := binarysubset[0]
	for _, b := range binarysubset {
		cumulativeAND = cumulativeAND & b
	}
	return cumulativeAND > 0
}
