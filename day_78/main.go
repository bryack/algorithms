package main

func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	if n > m {
		s, t = t, s
		m, n = n, m
	}

	for i := range s {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}
	return m-n == 1
}
