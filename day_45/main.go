package main

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	wState := 0
	res := math.MinInt

	for end := range nums {
		wState += nums[end]

		if end-begin+1 == k {
			res = max(res, wState)
			wState -= nums[begin]
			begin++
		}
	}
	return float64(res) / float64(k)
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	begin := 0
	wState := 0
	res := 0

	for end := range arr {
		wState += arr[end]
		if end-begin+1 == k {
			if wState/k >= threshold {
				res++
			}
			wState -= arr[begin]
			begin++
		}
	}

	return res
}

func countGoodSubstrings(s string) int {
	begin := 0
	res := 0

	for end := range s {
		if end-begin+1 == 3 {
			if s[begin] != s[begin+1] && s[begin+1] != s[end] && s[begin] != s[end] {
				res++
			}
			begin++
		}
	}
	return res
}

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

	begin := 0
	wState := 0

	for end := range nums {
		wState += nums[end]

		if end-begin == 2*k {
			res[begin+k] = wState / (end - begin + 1)
			wState -= nums[begin]
			begin++
		}
	}
	return res
}

func maximumSubarraySum(nums []int, k int) int64 {
	begin := 0
	m := make(map[int]int, k)
	wState := 0
	res := 0

	for end := range nums {
		wState += nums[end]
		m[nums[end]]++

		if end-begin+1 == k {
			if len(m) == k {
				res = max(res, wState)
			}
			m[nums[begin]]--
			if m[nums[begin]] <= 0 {
				delete(m, nums[begin])
			}
			wState -= nums[begin]
			begin++
		}
	}
	return int64(res)
}

func maxVowels(s string, k int) int {
	begin := 0
	wState := 0
	res := 0

	for end := range s {
		if isVowel(s[end]) {
			wState++
		}

		if end-begin+1 == k {
			res = max(res, wState)
			if isVowel(s[begin]) {
				wState--
			}
			begin++
		}
	}
	return res
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

func minimumRecolors(blocks string, k int) int {
	begin := 0
	wState := 0
	res := 101

	for end := range blocks {
		if blocks[end] == 'W' {
			wState++
		}
		if end-begin+1 == k {
			res = min(res, wState)
			if blocks[begin] == 'W' {
				wState--
			}
			begin++
		}
	}
	return res
}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	begin := 0
	var arr [101]int
	res := make([]int, 0, len(nums))

	for end := range nums {
		arr[nums[end]+50]++

		if end-begin+1 == k {
			count := 0
			found := false
			for i := 0; i < 50; i++ {
				count += arr[i]
				if count >= x {
					res = append(res, i-50)
					found = true
					break
				}
			}
			if !found {
				res = append(res, 0)
			}
			arr[nums[begin]+50]--
			begin++
		}
	}
	return res
}
