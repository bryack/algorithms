package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sub  string
		want []int
	}{
		{name: "test1", s: "cbaebabacd", sub: "abc", want: []int{0, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findAnagrams(tt.s, tt.sub)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test1", nums: []int{1, 0, 1, 1, 0}, want: 4},
		{name: "test2", nums: []int{1, 0, 1, 1, 0, 1}, want: 4},
		// single element
		{name: "single one", nums: []int{1}, want: 1},
		{name: "single zero", nums: []int{0}, want: 1},

		// all ones
		{name: "all ones", nums: []int{1, 1, 1, 1}, want: 4},

		// all zeros
		{name: "all zeros", nums: []int{0, 0, 0, 0}, want: 1},

		// zero in the middle
		{name: "bridge two groups", nums: []int{1, 1, 0, 1, 1}, want: 5},

		// zero at the beginning
		{name: "zero at beginning", nums: []int{0, 1, 1, 1}, want: 4},

		// zero at the end
		{name: "zero at end", nums: []int{1, 1, 1, 0}, want: 4},

		// two zeros
		{name: "two zeros", nums: []int{1, 1, 0, 1, 0, 1, 1}, want: 4},

		// long block
		{name: "long block", nums: []int{1, 1, 1, 0, 1, 1, 1}, want: 7},

		// best window not touching edges
		{name: "middle window", nums: []int{0, 1, 1, 0, 1, 1, 1, 0}, want: 6},

		// alternating
		{name: "alternating", nums: []int{1, 0, 1, 0, 1, 0, 1}, want: 3},

		// many zeros
		{name: "many zeros", nums: []int{0, 1, 0, 0, 1, 0}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxConsecutiveOnes(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "test1", s: "eceba", want: 3},
		{name: "example2", s: "ccaabbb", want: 5},

		// one character
		{name: "single char", s: "a", want: 1},

		// all same
		{name: "all same", s: "aaaaa", want: 5},

		// exactly two distinct
		{name: "two distinct", s: "abababab", want: 8},

		// three distinct
		{name: "three distinct", s: "abc", want: 2},

		// third char at the end
		{name: "third at end", s: "aabbc", want: 4},

		// third char at the beginning
		{name: "third at beginning", s: "caabb", want: 4},

		// best window in the middle
		{name: "middle window", s: "abcbbbbcccbdddadacb", want: 10},

		// alternating two chars
		{name: "alternating", s: "ababababab", want: 10},

		// shrink multiple times
		{name: "multiple shrinks", s: "abaccc", want: 4},

		// long suffix
		{name: "long suffix", s: "aabacbebebe", want: 6},

		// all unique
		{name: "all unique", s: "abcdef", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringTwoDistinct(tt.s)
			assert.Equal(t, tt.want, res)
		})
	}
}
