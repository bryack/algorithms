package main

import "slices"

func isSubsequence(sub string, s string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(sub) && s[i] == sub[j] {
			j++
		}
	}
	return j == len(sub)
}

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	res := 0
	i, j := 0, len(people)-1

	for i <= j {
		sum := people[i] + people[j]
		if sum <= limit {
			i++
			j--
		} else {
			j--
		}
		res++
	}
	return res
}
