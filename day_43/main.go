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
