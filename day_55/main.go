package main

import (
	"fmt"
	"strconv"
)

func trap(height []int) int {
	maxL, maxR := 0, 0
	i, j := 0, len(height)-1
	res := 0

	for i < j {

		if height[i] < height[j] {

			if height[i] >= maxL {
				maxL = height[i]
			} else {
				res += maxL - height[i]
			}
			i++
		} else {

			if height[j] >= maxR {
				maxR = height[j]
			} else {
				res += maxR - height[j]
			}
			j--
		}

	}
	return res
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
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
	res := []string{}
	end := 0

	for end < len(nums) {
		start := end
		for end < len(nums)-1 && nums[end]+1 == nums[end+1] {
			end++
		}
		count := end - start
		if count > 0 {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[end]))
		}
		end++
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
		length := end - start

		if length > 1 {
			count := strconv.Itoa(length)
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

	for end < len(s) {
		start := end

		for end < len(s) && buf[end] != ' ' {
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
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	buf := []byte{}

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
		buf = append(buf, byte(sum%10)+'0')
	}

	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}

	return string(buf)
}

func countSubstrings(s string) int {
	res := 0
	n := len(s)

	for i := 0; i < 2*n-1; i++ {
		l := i / 2
		r := l + i%2

		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
