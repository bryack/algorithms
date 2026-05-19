package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	w_state := 0
	res := math.MinInt

	for end := range nums {
		w_state += nums[end]
		w_size := end - begin + 1

		if w_size == k {
			res = max(res, w_state)
			w_state -= nums[begin]
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

	for end := range s {
		if end-begin+1 == 3 {
			if s[begin] != s[begin+1] && s[begin] != s[end] && s[end] != s[begin+1] {
				res++
			}
			begin++
		}
	}
	return res
}

func getAverages(nums []int, k int) []int {
	begin := 0
	w_state := 0
	res := make([]int, len(nums))

	for i := range res {
		res[i] = -1
	}

	for end := range nums {
		w_state += nums[end]
		if end-begin == 2*k {
			res[begin+k] = w_state / (end - begin + 1)
			w_state -= nums[begin]
			begin++
		}
	}
	return res
}
