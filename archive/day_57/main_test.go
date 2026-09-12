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

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{name: "test1", s1: "ab", s2: "eidbaooo", want: true},
		{name: "test2", s1: "ab", s2: "eidboaoo", want: false},
		{name: "s1 longer than s2", s1: "abc", s2: "ab", want: false},
		{name: "identical strings", s1: "abc", s2: "abc", want: true},
		{name: "permutation at start", s1: "abc", s2: "abcde", want: true},
		{name: "permutation at end", s1: "abc", s2: "defabc", want: true},
		{name: "not enough characters", s1: "abc", s2: "aabbc", want: false},
		{name: "single char found", s1: "a", s2: "bca", want: true},
		{name: "single char not found", s1: "z", s2: "abc", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := checkInclusion(tt.s1, tt.s2)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		sub  string
		want []int
	}{
		{name: "test1", s1: "cbaebabacd", sub: "abc", want: []int{0, 6}},
		{name: "test2", s1: "abab", sub: "ab", want: []int{0, 1, 2}},
		{name: "no anagrams", s1: "abcde", sub: "xyz", want: []int{}},
		{name: "full string match", s1: "abc", sub: "abc", want: []int{0}},
		{name: "sub longer than s", s1: "a", sub: "abc", want: []int{}},
		{name: "overlapping anagrams", s1: "aaaaa", sub: "aa", want: []int{0, 1, 2, 3}},
		{name: "single char anagram", s1: "abac", sub: "a", want: []int{0, 2}},
		{name: "all same", s1: "aaaa", sub: "aa", want: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findAnagrams(tt.s1, tt.sub)
			assert.Equal(t, tt.want, res)
		})
	}
}
