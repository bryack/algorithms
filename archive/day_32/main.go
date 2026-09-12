package main

import (
	"math"
	"unicode"
)

func maxNumberOfBalloons(text string) int {
	if len(text) < len("balloon") {
		return 0
	}
	var b [26]int
	var t [26]int
	balloon := "balloon"
	count := math.MaxInt

	for _, char := range balloon {
		b[char-'a']++
	}

	for _, char := range text {
		t[char-'a']++
	}

	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			continue
		}
		count = min(count, t[i]/b[i])
	}

	if count == math.MaxInt {
		count = 0
	}

	return count
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	i, j := 0, len(height)-1
	maxLeft, maxRight := height[0], height[len(height)-1]
	count := 0

	for i < j {
		if maxLeft <= maxRight {
			i++
			if height[i] >= maxLeft {
				maxLeft = max(maxLeft, height[i])
			} else {
				count += maxLeft - height[i]
			}
		} else {
			j--
			if height[j] >= maxRight {
				maxRight = max(maxRight, height[j])
			} else {
				count += maxRight - height[j]
			}
		}
	}
	return count
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
