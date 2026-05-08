package main

import (
	"unicode"
)

func topKFrequent(nums []int, k int) []int {
	s := make([]int, 0, k)
	m := make(map[int]int, len(nums))

	for _, n := range nums {
		m[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, v := range m {
		bucket[v] = append(bucket[v], key)
	}

	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		for j := 0; j < len(bucket[i]); j++ {
			if len(s) == k {
				return s
			}
			s = append(s, bucket[i][j])
		}
	}
	return s
}

func reverseOnlyLetters(s string) string {
	reverse := make([]byte, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		if !unicode.IsLetter(rune(s[i])) {
			reverse[i] = byte(s[i])
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) {
			reverse[j] = byte(s[j])
			j--
			continue
		}
		reverse[i] = byte(s[j])
		reverse[j] = byte(s[i])
		i++
		j--
	}
	return string(reverse)
}
