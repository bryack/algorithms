package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0, len(nums)/2)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func getAverages(nums []int, k int) []int {
	begin := 0
	wState := 0
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

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
	wState := 0
	res := 0
	m := map[int]int{}

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
	count := 0
	maxCount := 0

	for end := range s {
		if isVowel(s[end]) {
			count++
		}

		if end-begin+1 == k {
			maxCount = max(maxCount, count)
			if isVowel(s[begin]) {
				count--
			}
			begin++
		}
	}
	return maxCount
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}
