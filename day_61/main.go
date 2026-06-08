package main

import "slices"

func reverseOnlyLetters(s string) string {
	buf := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		if !isLetter(buf[i]) {
			i++
			continue
		}
		if !isLetter(buf[j]) {
			j--
			continue
		}

		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func minMeetingRooms(intervals [][]int) int {
	first := make([]int, len(intervals))
	second := make([]int, len(intervals))

	for i := range intervals {
		first[i] = intervals[i][0]
		second[i] = intervals[i][1]
	}

	slices.Sort(first)
	slices.Sort(second)

	res := 0
	end := 0

	for start := 0; start < len(intervals); start++ {
		if first[start] < second[end] {
			res++
		} else {
			end++
		}
	}
	return res
}
