package main

import (
	"math"
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i <= j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	window_state := 0
	res := math.MinInt

	for end := range nums {
		window_state += nums[end]

		if end-begin+1 == k {
			res = max(res, window_state)
			window_state -= nums[begin]
			begin++
		}
	}
	return float64(res) / float64(k)
}
