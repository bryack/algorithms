package main

import "math"

// [1, 2, 3, 4, 5] k = 3 -- 12
// [-10, -5, -2, 1, 3, 5, 20] k = 4 -- 29
// [3, 1, 5, -6, 5] k = 1 -- 5
// [0, 0, 0, 0, 0] k = 0 -- 0
// [7, 7, 7, 7, 7] k = 5 -- 0
// [] k = -1 -- 0
// [1, 2] k = 0 -- 0

func maxSumSubArray(nums []int, k int) int {
	wState, res, begin := 0, math.MinInt, 0
	m := map[int]int{}

	for end := range nums {
		wState += nums[end]
		m[nums[end]]++

		if end-begin+1 == k {
			if len(m) == k {
				res = max(wState, res)
			}
			m[nums[begin]]--
			if m[nums[begin]] == 0 {
				delete(m, nums[begin])
			}
			wState -= nums[begin]
			begin++
		}
	}
	if res == math.MinInt {
		return 0
	}
	return res
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	first, second := [2001]bool{}, [2001]bool{}

	for _, n := range nums1 {
		first[n+1000] = true
	}
	for _, n := range nums2 {
		second[n+1000] = true
	}

	res := make([][]int, 2)

	for i := 0; i < 2001; i++ {
		if first[i] && !second[i] {
			res[0] = append(res[0], i-1000)
		}
		if !first[i] && second[i] {
			res[1] = append(res[1], i-1000)
		}
	}
	return res
}
