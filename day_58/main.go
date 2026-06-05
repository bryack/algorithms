package main

import (
	"fmt"
	"slices"
	"strconv"
)

func longestPalindrome(s string) string {
	maxL, maxR := 0, 0
	n := len(s)

	for center := 0; center < 2*n-1; center++ {
		l := center / 2
		r := l + center%2

		for l >= 0 && r < n && s[l] == s[r] {
			if maxR-maxL < r-l {
				maxL, maxR = l, r
			}

			l--
			r++
		}
	}
	return s[maxL : maxR+1]
}

func summaryRanges(nums []int) []string {
	i := 0
	res := []string{}

	for i < len(nums) {
		j := i

		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		if nums[i] == nums[j] {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[j], nums[i]))
		}
		i++
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
			for i := range count {
				chars[idx] = count[i]
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

	for end := 0; end < len(buf); end++ {
		start := end

		for end < len(buf) && buf[end] != ' ' {
			end++
		}

		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			buf[i], buf[j] = buf[j], buf[i]
		}
	}
	return string(buf)
}

func addStrings(num1 string, num2 string) string {
	buf := []byte{}
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
		buf = append(buf, byte(sum%10)+'0')
	}

	for x, y := 0, len(buf)-1; x < y; x, y = x+1, y-1 {
		buf[x], buf[y] = buf[y], buf[x]
	}

	return string(buf)

}

func countSubstrings(s string) int {
	res := 0
	n := len(s)

	for center := 0; center < 2*n-1; center++ {
		l := center / 2
		r := l + center%2

		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}

func removeElement(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	res := []int{}

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(res) == 0 || res[len(res)-1] != nums1[i] {
				res = append(res, nums1[i])
			}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	res := []int{}

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func removeDuplicates(nums []int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if j < 2 || nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
