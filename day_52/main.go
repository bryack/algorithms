package main

import (
	"fmt"
	"slices"
)

func isSubsequence(sub string, s string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(sub) && s[i] == sub[j] {
			j++
		}
	}
	return j == len(sub)
}

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	res := 0
	i, j := 0, len(people)-1

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
	res := make([]string, 0, len(nums))
	i, j := 0, 0

	for i < len(nums) {
		for i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			i++
		}

		if nums[j] == nums[i] {
			res = append(res, fmt.Sprintf("%d", nums[j]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[j], nums[i]))
		}
		i++
		j = i

	}
	return res
}

func compress(chars []byte) int {
	idx := 0
	end := 0

	writeCount := func(c int, idx int) int {
		temp := idx
		for c != 0 {
			digit := c % 10
			chars[idx] = byte('0' + digit)
			c /= 10
			idx++
		}
		start := temp
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
