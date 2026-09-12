package main

func countSubstrings(s string) int {
	n := len(s)
	res := 0

	for center := 0; center < 2*n-1; center++ {
		l := center / 2
		r := l + center%2

		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
