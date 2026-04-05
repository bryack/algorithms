package main

import (
	"slices"
)

func findGCD(nums []int) int {
	minNum := slices.Min(nums)
	maxNum := slices.Max(nums)

	if maxNum%minNum == 0 {
		return minNum
	}

	for {
		d := maxNum % minNum
		if d == 0 {
			return minNum
		}
		maxNum = minNum
		minNum = d
	}
}
