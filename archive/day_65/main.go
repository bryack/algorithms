package main

import (
	"strconv"
)

func compress(chars []byte) int {
	idx := 0
	end := 0

	for end < len(chars) {
		start := end

		for end < len(chars) && chars[start] == chars[end] {
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

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	begin := 0
	res := 0
	m := map[byte]int{}

	for end := range s {
		m[s[end]]++

		for len(m) > k {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}
