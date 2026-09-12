package main

import (
	"slices"
)

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

func minimumRecolors(blocks string, k int) int {
	begin := 0
	count := 0
	minRecolors := 100

	for end := range blocks {
		if blocks[end] == 'W' {
			count++
		}

		if end-begin+1 == k {
			minRecolors = min(minRecolors, count)
			if blocks[begin] == 'W' {
				count--
			}
			begin++
		}
	}
	return minRecolors
}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	begin := 0
	var freq [101]int
	res := make([]int, 0, len(nums))

	for end := range nums {
		freq[nums[end]+50]++

		if end-begin+1 == k {
			count := 0
			found := false
			for i := 0; i < 50; i++ {
				count += freq[i]
				if count >= x {
					res = append(res, i-50)
					found = true
					break
				}
			}
			if !found {
				res = append(res, 0)
			}
			freq[nums[begin]+50]--
			begin++
		}
	}
	return res
}

func maxScore(cardPoints []int, k int) int {
	i, j := 0, len(cardPoints)-k
	wState := 0

	for x := j; x < len(cardPoints); x++ {
		wState += cardPoints[x]
	}
	res := wState

	for j < len(cardPoints) {
		wState += (cardPoints[i] - cardPoints[j])
		res = max(res, wState)
		i++
		j++
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	begin := 0
	wState := 0
	res := len(nums) + 1

	for end := range nums {
		wState += nums[end]

		for wState >= target {
			wSize := end - begin + 1
			res = min(res, wSize)
			wState -= nums[begin]
			begin++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i <= j {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
