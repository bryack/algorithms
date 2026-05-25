package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	k = k % len(nums)

	swap := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	swap(0, len(nums)-1)
	swap(0, k-1)
	swap(k, len(nums)-1)
}

func maxArea(height []int) int {
	maxArea := 0

	i, j := 0, len(height)-1

	for i < j {
		if height[i] < height[j] {
			cArea := height[i] * (j - i)
			maxArea = max(maxArea, cArea)
			i++
		} else {
			cArea := height[j] * (j - i)
			maxArea = max(maxArea, cArea)
			j--
		}
	}
	return maxArea
}

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	i, j := 0, len(people)-1
	res := 0

	for i <= j {
		sum := people[i] + people[j]
		if sum > limit {
			j--
		} else {
			i++
			j--
		}
		res++
	}
	return res
}

func mergeAlternately(word1 string, word2 string) string {
	buf := make([]byte, len(word1)+len(word2))
	i, j := 0, 0
	idx := 0

	for i < len(word1) && j < len(word2) {
		buf[idx] = word1[i]
		i++
		idx++
		buf[idx] = word2[j]
		j++
		idx++
	}

	for i < len(word1) {
		buf[idx] = word1[i]
		i++
		idx++
	}
	for j < len(word2) {
		buf[idx] = word2[j]
		j++
		idx++
	}
	return string(buf)
}

func trap(height []int) int {
	maxL, maxR := 0, 0
	i, j := 0, len(height)-1
	res := 0

	for i < j {
		if maxL > height[i] {
			res += maxL - height[i]
		}
		if maxR > height[j] {
			res += maxR - height[j]
		}
		maxL = max(maxL, height[i])
		maxR = max(maxR, height[j])

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func reverseOnlyLetters(s string) string {
	buf := []byte(s)
	i, j := 0, len(s)-1

	for i <= j {
		if !isLetter(buf[i]) {
			i++
			continue
		}
		if !isLetter(buf[j]) {
			j--
			continue
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func longestPalindrome(s string) string {
	l, r := 0, 0

	expand := func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			if r-l < j-i {
				l, r = i, j
			}
			i--
			j++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[l : r+1]
}

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = nextValid(s, i)
		j = nextValid(t, j)

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

func nextValid(s string, i int) int {
	count := 0
	for i >= 0 {
		if s[i] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			return i
		}
		i--
	}
	return -1
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0, len(nums))

	for end := 0; end < len(nums); end++ {
		start := end

		for end < len(nums)-1 && nums[end+1] == nums[end]+1 {
			end++
		}
		if end-start > 0 {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[end]))
		}
	}
	return res
}

func compress(chars []byte) int {
	read := 0
	write := 0

	for read < len(chars) {
		start := read
		for read < len(chars) && chars[read] == chars[start] {
			read++
		}
		count := read - start
		chars[write] = chars[start]
		write++
		if count > 1 {
			write = writeCount(chars, write, count)
		}
	}
	return write
}

func writeCount(chars []byte, write, count int) int {
	i := write
	for count != 0 {
		digit := count % 10
		chars[write] = byte('0' + digit)
		count /= 10
		write++
	}
	start := i
	end := write - 1

	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}

	return write
}
