package main

import (
	"fmt"
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
