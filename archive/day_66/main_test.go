package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneEditDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "insert middle", s: "ab", t: "acb", want: true},
		{name: "insert end", s: "ab", t: "abc", want: true},
		{name: "insert start", s: "ab", t: "cab", want: true},
		{name: "delete middle", s: "acb", t: "ab", want: true},
		{name: "delete end", s: "abc", t: "ab", want: true},
		{name: "delete start", s: "cab", t: "ab", want: true},
		{name: "replace", s: "ab", t: "ac", want: true},
		{name: "replace single", s: "1203", t: "1213", want: true},
		{name: "empty to single", s: "", t: "A", want: true},
		{name: "single to empty", s: "A", t: "", want: true},
		{name: "identical", s: "ab", t: "ab", want: false},
		{name: "both empty", s: "", t: "", want: false},
		{name: "diff > 1 length", s: "ab", t: "abcb", want: false},
		{name: "diff > 1 length reverse", s: "abcb", t: "ab", want: false},
		{name: "two replacements", s: "ab", t: "cd", want: false},
		{name: "two inserts needed", s: "a", t: "abc", want: false},
		{name: "two deletes needed", s: "abc", t: "a", want: false},
		{name: "single replace char", s: "a", t: "b", want: true},
		{name: "single replace case", s: "a", t: "A", want: true},
		{name: "transposition", s: "ab", t: "ba", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isOneEditDistance(tt.s, tt.t)
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
