package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	i, j := 0, 0
	res := []string{}

	for i < len(nums) {

		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		if nums[i] == nums[j] {
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
	end := 0
	idx := 0

	for end < len(chars) {
		start := end

		for end < len(chars) && chars[start] == chars[end] {
			end++
		}

		chars[idx] = chars[start]
		idx++

		count := end - start
		if count > 1 {
			num := strconv.Itoa(count)
			for _, n := range num {
				chars[idx] = byte(n)
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

	for i := 0; i < m; i++ {
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
	res := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		li := max(firstList[i][0], secondList[j][0])
		ri := min(firstList[i][1], secondList[j][1])

		if li <= ri {
			res = append(res, []int{li, ri})
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

		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			buf[i], buf[j] = buf[j], buf[i]
		}

		end++

	}
	return string(buf)
}

func addStrings(num1 string, num2 string) string {
	res := []byte{}
	carry := 0
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		carry = sum / 10
		res = append(res, byte(sum%10)+'0')
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func countSubstrings(s string) int {
	res := 0
	n := len(s)

	for center := 0; center < 2*n+1; center++ {
		left := center / 2
		right := left + center%2

		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}
