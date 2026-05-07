package main

import "math"

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
