package main

func trap(height []int) int {
	maxL, maxR := 0, 0
	res := 0
	i, j := 0, len(height)-1

	for i < j {
		if height[i] < height[j] {
			if maxL < height[i] {
				maxL = height[i]
			} else {
				res += maxL - height[i]
			}
			i++
		} else {
			if maxR < height[j] {
				maxR = height[j]
			} else {
				res += maxR - height[j]
			}
			j--
		}
	}
	return res
}

// 1234 -> 1254 замена
// ab  -> acb   вставка
// ab   -> abc  удаление
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
