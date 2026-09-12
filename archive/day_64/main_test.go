package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "test1", s: "eceba", k: 2, want: 3},
		{name: "example2", s: "aa", k: 1, want: 2},

		// edge cases
		{name: "empty k", s: "abc", k: 0, want: 0},
		{name: "single char", s: "a", k: 1, want: 1},
		{name: "single char k zero", s: "a", k: 0, want: 0},

		// all same
		{name: "all same", s: "aaaaaa", k: 1, want: 6},
		{name: "all same large k", s: "aaaaaa", k: 5, want: 6},

		// exactly k distinct
		{name: "exactly k distinct", s: "ababab", k: 2, want: 6},

		// all unique
		{name: "all unique k one", s: "abcdef", k: 1, want: 1},
		{name: "all unique k two", s: "abcdef", k: 2, want: 2},
		{name: "all unique k three", s: "abcdef", k: 3, want: 3},

		// k larger than distinct count
		{name: "k larger than distinct", s: "abc", k: 10, want: 3},

		// shrink window
		{name: "shrink once", s: "abac", k: 2, want: 3},
		{name: "shrink many times", s: "abaccc", k: 2, want: 4},

		// optimal window in middle
		{name: "middle window", s: "aabacbebebe", k: 3, want: 7},

		// frequent character changes
		{name: "alternating", s: "abcabcabc", k: 2, want: 2},

		// long valid suffix
		{name: "long suffix", s: "abcbbbbcccbdddadacb", k: 2, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct(tt.s, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}
