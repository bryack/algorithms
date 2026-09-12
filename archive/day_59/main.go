package main

import (
	"slices"
	"strconv"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	i, j := 0, len(people)-1
	res := 0

	for i <= j {
		sum := people[i] + people[j]
		if sum <= limit {
			i++
			j--
		} else {
			j--
		}
		res++
	}
	return res
}

func compress(chars []byte) int {
	idx := 0
	end := 0

	for end < len(chars) {
		start := end

		for end < len(chars) && chars[end] == chars[start] {
			end++
		}

		chars[idx] = chars[start]
		idx++

		if end-start > 1 {
			count := strconv.Itoa(end - start)
			for _, digit := range count {
				chars[idx] = byte(digit)
				idx++
			}
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

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}
	return n-m == 1
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	firstChar := needle[0]

	for i := range len(haystack) + 1 - len(needle) {
		if haystack[i] == firstChar {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	start := make([]int, n)
	end := make([]int, n)

	for i := range intervals {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}

	slices.Sort(start)
	slices.Sort(end)

	res := 0
	endP := 0

	for startP := 0; startP < n; startP++ {
		if start[startP] < end[endP] {
			res++
		} else {
			endP++
		}
	}
	return res
}

func maximumSubarraySum(nums []int, k int) int64 {
	res := 0
	begin := 0
	wState := 0
	m := map[int]int{}

	for end := range nums {
		wState += nums[end]
		m[nums[end]]++

		if end-begin+1 == k {
			if len(m) == k {
				res = max(res, wState)
			}
			m[nums[begin]]--
			if m[nums[begin]] < 1 {
				delete(m, nums[begin])
			}
			wState -= nums[begin]
			begin++
		}
	}
	return int64(res)
}
