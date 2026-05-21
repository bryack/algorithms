package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfSubarrays(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		k         int
		threshold int
		want      int
	}{
		{name: "test1", arr: []int{2, 2, 2, 2, 5, 5, 5, 8}, k: 3, threshold: 4, want: 3},
		{name: "test2", arr: []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, k: 3, threshold: 5, want: 6},
		{name: "len_eq_k", arr: []int{1, 2, 3}, k: 3, threshold: 2, want: 1},
		{name: "all_below_threshold", arr: []int{1, 1, 1}, k: 2, threshold: 2, want: 0},
		{name: "k_eq_1", arr: []int{5, 1, 5}, k: 1, threshold: 5, want: 2},
		{name: "all_equal_threshold", arr: []int{3, 3, 3}, k: 2, threshold: 3, want: 2},
		{name: "single_valid", arr: []int{1, 4, 1}, k: 1, threshold: 4, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numOfSubarrays(tt.arr, tt.k, tt.threshold)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestCountGoodSubstrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "test1", input: "xyzzaz", want: 1},
		{name: "test2", input: "aababcabc", want: 4},
		{name: "empty", input: "", want: 0},
		{name: "len_2", input: "ab", want: 0},
		{name: "len_3_good", input: "abc", want: 1},
		{name: "len_3_bad", input: "aaa", want: 0},
		{name: "all_distinct", input: "abcdef", want: 4},
		{name: "all_same", input: "aaaa", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countGoodSubstrings(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGetAverages(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "test1", nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, k: 3, want: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
		{name: "test2", nums: []int{100000}, k: 0, want: []int{100000}},
		{name: "test3", nums: []int{8}, k: 100000, want: []int{-1}},
		{name: "n3_k1", nums: []int{1, 2, 3}, k: 1, want: []int{-1, 2, -1}},
		{name: "all_same", nums: []int{5, 5, 5, 5, 5}, k: 1, want: []int{-1, 5, 5, 5, -1}},
		{name: "n2_k1", nums: []int{1, 2}, k: 1, want: []int{-1, -1}},
		{name: "negatives", nums: []int{-1, -2, -3, -4, -5}, k: 1, want: []int{-1, -2, -3, -4, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getAverages(tt.nums, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}

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

func TestGetSubarrayBeauty(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		x    int
		want []int
	}{
		{name: "classic", nums: []int{1, -1, -3, -2, 3}, k: 3, x: 2, want: []int{-1, -2, -2}},
		{name: "all negative x=2", nums: []int{-1, -2, -3, -4, -5}, k: 2, x: 2, want: []int{-1, -2, -3, -4}},
		{name: "mixed x=1", nums: []int{-3, 1, 2, -3, 0, -3}, k: 2, x: 1, want: []int{-3, 0, -3, -3, -3}},
		{name: "all positive", nums: []int{1, 2, 3, 4}, k: 2, x: 1, want: []int{0, 0, 0}},
		{name: "all negative x=1", nums: []int{-5, -4, -3, -2}, k: 2, x: 1, want: []int{-5, -4, -3}},
		{name: "k equals len", nums: []int{1, -2, 3, -4}, k: 4, x: 2, want: []int{-2}},
		{name: "x equals k", nums: []int{-1, -2, -3}, k: 3, x: 3, want: []int{-1}},
		{name: "positive xth element", nums: []int{-5, 1, 2}, k: 3, x: 2, want: []int{0}},
		{name: "single window", nums: []int{-1}, k: 1, x: 1, want: []int{-1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getSubarrayBeauty(tt.nums, tt.k, tt.x)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name       string
		cardPoints []int
		k          int
		want       int
	}{
		{name: "test1", cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3, want: 12},
		{name: "test2", cardPoints: []int{2, 2, 2}, k: 2, want: 4},
		{name: "test3", cardPoints: []int{9, 7, 7, 9, 7, 7, 9}, k: 7, want: 55},
		{name: "test4", cardPoints: []int{1, 100, 1, 1, 1, 1, 1}, k: 3, want: 102},
		{name: "test5", cardPoints: []int{1, 1, 1, 1, 1, 100, 1}, k: 3, want: 102},
		{name: "test6", cardPoints: []int{5}, k: 1, want: 5},
		{name: "test7", cardPoints: []int{1, 2, 3, 4, 5}, k: 1, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxScore(tt.cardPoints, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test1", nums: []int{1, 1, 0, 1}, want: 3},
		{name: "test2", nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, want: 5},
		{name: "test3", nums: []int{1, 1, 1}, want: 2},
		{name: "all_zeros", nums: []int{0, 0, 0}, want: 0},
		{name: "multiple_zeros", nums: []int{1, 0, 0, 0, 1}, want: 1},
		{name: "zero_at_end", nums: []int{1, 1, 1, 0}, want: 3},
		{name: "zero_at_start", nums: []int{0, 1, 1, 1}, want: 3},
		{name: "single_zero", nums: []int{0}, want: 0},
		{name: "single_one", nums: []int{1}, want: 0},
		{name: "alternating", nums: []int{1, 0, 1, 0, 1, 0, 1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubarray(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}
