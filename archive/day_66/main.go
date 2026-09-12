package main

func isOneEditDistance(s string, t string) bool {
	m, n := len(s), len(t)

	if m > n {
		s, t = t, s
		m, n = n, m
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

func minWindow(s string, sub string) string {
	var need [128]int
	var window [128]int

	for i := 0; i < len(sub); i++ {
		need[sub[i]]++
	}

	required := len(sub)
	satisfied := 0
	start := 0
	minLen := len(s) + 1
	minLeft := 0

	for end := 0; end < len(s); end++ {
		window[s[end]]++

		if need[s[end]] > 0 && window[s[end]] <= need[s[end]] {
			satisfied++
		}

		for satisfied == required {
			if length := end - start + 1; length < minLen {
				minLen = length
				minLeft = start
			}
			window[s[start]]--

			if need[s[start]] > 0 && window[s[start]] < need[s[start]] {
				satisfied--
			}
			start++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	res := 0
	begin := 0
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
