package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = getNextValid(s, i)
		j = getNextValid(t, j)

		if i < 0 && j < 0 {
			return true
		}

		if i < 0 || j < 0 {
			return false
		}

		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func getNextValid(s string, idx int) int {
	count := 0
	for idx >= 0 {
		if s[idx] == '#' {
			count++
			idx--
		} else if count > 0 {
			count--
			idx--
		} else {
			return idx
		}
	}
	return -1
}

func summaryRanges(nums []int) []string {
	res := []string{}
	end := 0

	for end < len(nums) {
		start := end

		for end < len(nums)-1 && nums[end]+1 == nums[end+1] {
			end++
		}

		if nums[start] == nums[end] {
			res = append(res, fmt.Sprintf("%d", nums[end]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		}
		end++
	}
	return res
}

func removeDuplicates(nums []int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0

	for i < j {
		curA := (min(height[j], height[i])) * (j - i)
		res = max(res, curA)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func compress(chars []byte) int {
	idx := 0
	end := 0

	writeCount := func(count, idx int) int {
		i := idx
		for count != 0 {
			digit := count % 10
			chars[idx] = byte('0' + digit)
			count /= 10
			idx++
		}

		start := i
		end := idx - 1

		for start < end {
			chars[start], chars[end] = chars[end], chars[start]
			start++
			end--
		}
		return idx
	}

	for end < len(chars) {
		start := end

		for end < len(chars) && chars[start] == chars[end] {
			end++
		}

		count := end - start
		chars[idx] = chars[start]
		idx++
		if count > 1 {
			idx = writeCount(count, idx)
		}
	}
	return idx
}

func isOneEditDistance(s string, t string) bool {
	m, n := len(s), len(t)

	if m > n {
		s, t = t, s
		m, n = n, m
	}

	if n-m > 1 {
		return false
	}

	for i := 0; i < m; i++ {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}
	return n-m == 1
}
