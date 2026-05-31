package main

import (
	"fmt"
	"strconv"
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

		if nums[end] == nums[start] {
			res = append(res, fmt.Sprintf("%d", nums[start]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		}
		end++
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

		count := end - start
		if count > 1 {
			temp := strconv.Itoa(count)
			for _, digit := range temp {
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

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	res := [][]int{}

	for i < len(firstList) && j < len(secondList) {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])

		if lo <= hi {
			res = append(res, []int{lo, hi})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func reverseWords(s string) string {
	buf := []byte(s)
	end := 0

	for end < len(buf) {
		start := end
		for end < len(buf) && buf[end] != ' ' {
			end++
		}
		if start == end {
			end++
		} else {
			reverse(buf, start, end-1)
		}

	}
	return string(buf)
}

func reverse(buf []byte, i, j int) {
	for i <= j {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
}
