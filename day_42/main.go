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

func numOfSubarrays(arr []int, k int, threshold int) int {
	begin := 0
	w_state := 0
	res := 0

	for end := range arr {
		w_state += arr[end]
		w_size := end - begin + 1

		if w_size == k {
			if w_state/k >= threshold {
				res++
			}
			w_state -= arr[begin]
			begin++
		}
	}
	return res
}

func countGoodSubstrings(s string) int {
	begin := 0
	res := 0

	for end := 0; end < len(s); end++ {
		if end-begin == 2 {
			if isGood(s[begin], s[begin+1], s[end]) {
				res++
			}
			begin++
		}
	}
	return res
}

func isGood(begin, center, end byte) bool {
	return begin != center && begin != end && center != end
}
