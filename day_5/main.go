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

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m := make(map[int]int)
	for _, card := range deck {
		m[card]++
	}
	s := make([]int, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}

	d := s[0]
	for i := 0; i < len(s)-1; i++ {
		d = GCD(d, s[i+1])
	}

	return d >= 2
}

func GCD(a, b int) int {
	x, y := max(a, b), min(a, b)
	for {
		d := x % y
		if d == 0 {
			return y
		}
		x = y
		y = d
	}
}
