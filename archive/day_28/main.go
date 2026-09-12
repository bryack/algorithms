package main

import "strings"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums)/2)

	for i, n := range nums {
		if v, ok := m[n]; ok {
			if i-v <= k {
				return true
			}
			m[n] = i
		} else {
			m[n] = i
		}
	}
	return false
}

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			if isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1) {
				return true
			} else {
				return false
			}
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	builder.Grow(len(word1) + len(word2))
	i := 0

	for i < len(word1) {
		if i < len(word2) {
			builder.WriteByte(word1[i])
			builder.WriteByte(word2[i])
		} else {
			builder.WriteByte(word1[i])
		}
		i++
	}
	if i < len(word2) {
		for i < len(word2) {
			builder.WriteByte(word2[i])
			i++
		}

	}
	return builder.String()
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	first, second, writeIdx := m-1, n-1, m+n-1

	for first >= 0 && second >= 0 {
		if nums1[first] >= nums2[second] {
			nums1[writeIdx] = nums1[first]
			first--
		} else {
			nums1[writeIdx] = nums2[second]
			second--
		}
		writeIdx--
	}

	for second >= 0 {
		nums1[writeIdx] = nums2[second]
		writeIdx--
		second--
	}
}
