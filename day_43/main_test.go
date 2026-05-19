package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{name: "classic valid", nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3, want: 15},
		{name: "all duplicates", nums: []int{4, 4, 4}, k: 3, want: 0},
		{name: "duplicates not adjacent", nums: []int{1, 2, 1, 2}, k: 3, want: 0},
		{name: "k equals 1", nums: []int{5, 1, 3}, k: 1, want: 5},
		{name: "len equals k all distinct", nums: []int{1, 2, 3}, k: 3, want: 6},
		{name: "len equals k duplicates not adjacent", nums: []int{1, 2, 1}, k: 3, want: 0},
		{name: "single valid window", nums: []int{9, 9, 1, 2, 3}, k: 3, want: 12},
		{name: "test8", nums: []int{5, 3, 3, 1, 1}, k: 3, want: 0},
		{name: "no valid window with stale map entry", nums: []int{1, 2, 3, 2, 10}, k: 4, want: 0},
		{name: "stale element creates false valid", nums: []int{1, 2, 3, 2, 4, 5}, k: 4, want: 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximumSubarraySum(tt.nums, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "classic", s: "abciiidef", k: 3, want: 3},
		{name: "all vowels", s: "aeiou", k: 2, want: 2},
		{name: "mixed", s: "leetcode", k: 3, want: 2},
		{name: "no vowels", s: "bcdfg", k: 3, want: 0},
		{name: "all same vowel", s: "aaaaa", k: 3, want: 3},
		{name: "k equals 1 vowel", s: "abcde", k: 1, want: 1},
		{name: "k equals 1 consonant", s: "bcdfg", k: 1, want: 0},
		{name: "k equals len all vowels", s: "aeiou", k: 5, want: 5},
		{name: "k equals len no vowels", s: "bcdfg", k: 5, want: 0},
		{name: "single char vowel", s: "a", k: 1, want: 1},
		{name: "single char consonant", s: "b", k: 1, want: 0},
		{name: "vowels at edges", s: "aeibcdou", k: 4, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxVowels(tt.s, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}
