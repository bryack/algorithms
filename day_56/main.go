package main

func trap(height []int) int {
	maxL, maxR := 0, 0
	i, j := 0, len(height)-1
	res := 0

	for i < j {
		if height[i] < height[j] {
			if height[i] > maxL {
				maxL = height[i]
			} else {
				res += maxL - height[i]
			}
			i++
		} else {
			if height[j] > maxR {
				maxR = height[j]
			} else {
				res += maxR - height[j]
			}
			j--
		}
	}
	return res
}

func longestPalindrome(s string) string {
	n := len(s)
	maxL, maxR := 0, 0

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

func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height)-1

	for i < j {
		current := min(height[i], height[j]) * (j - i)
		res = max(res, current)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
