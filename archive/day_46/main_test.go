package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{name: "equal_length", word1: "abc", word2: "pqr", want: "apbqcr"},
		{name: "word2_longer", word1: "ab", word2: "pqrs", want: "apbqrs"},
		{name: "word1_longer", word1: "abcd", word2: "pq", want: "apbqcd"},
		{name: "both_empty", word1: "", word2: "", want: ""},
		{name: "empty_word1", word1: "", word2: "abc", want: "abc"},
		{name: "empty_word2", word1: "abc", word2: "", want: "abc"},
		{name: "single_char_each", word1: "a", word2: "b", want: "ab"},
		{name: "single_vs_long", word1: "a", word2: "bcdef", want: "abcdef"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately(tt.word1, tt.word2)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{name: "classic", nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{name: "all_same", nums: []int{2, 2, 2, 2, 2}, target: 8, want: [][]int{{2, 2, 2, 2}}},
		{name: "empty", nums: []int{}, target: 0, want: [][]int{}},
		{name: "too_short", nums: []int{1, 2, 3}, target: 6, want: [][]int{}},
		{name: "no_match", nums: []int{1, 2, 3, 4}, target: 100, want: [][]int{}},
		{name: "all_zeros", nums: []int{0, 0, 0, 0}, target: 0, want: [][]int{{0, 0, 0, 0}}},
		{name: "negative_target", nums: []int{-3, -2, -1, 0, 0, 1, 2, 3}, target: -6, want: [][]int{{-3, -2, -1, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fourSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{name: "test 1", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
		{name: "test 3", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 7, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test 4", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 0, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test 5", input: []int{1}, k: 3, want: []int{1}},
		{name: "k_greater_than_len", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 10, want: []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input, tt.k)
			assert.Equal(t, tt.want, tt.input)
		})
	}
}

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  int
	}{
		{name: "test 1", input: []int{1, 2}, k: 3, want: 1},
		{name: "test 2", input: []int{3, 2, 2, 1}, k: 3, want: 3},
		{name: "test 3", input: []int{3, 5, 3, 4}, k: 5, want: 4},
		{name: "single_person", input: []int{1}, k: 5, want: 1},
		{name: "two_fit", input: []int{1, 2}, k: 4, want: 1},
		{name: "two_no_fit", input: []int{3, 3}, k: 5, want: 2},
		{name: "all_same_fit", input: []int{2, 2, 2, 2}, k: 4, want: 2},
		{name: "all_same_no_fit", input: []int{3, 3, 3}, k: 5, want: 3},
		{name: "exact_limit", input: []int{3, 2, 2, 1}, k: 5, want: 2},
		{name: "five_people_max_2_per_boat", input: []int{1, 1, 1, 1, 1}, k: 5, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numRescueBoats(tt.input, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "classic", input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "second_example", input: []int{4, 2, 0, 3, 2, 5}, want: 9},
		{name: "single", input: []int{5}, want: 0},
		{name: "two_elements", input: []int{5, 4}, want: 0},
		{name: "ascending", input: []int{1, 2, 3, 4, 5}, want: 0},
		{name: "descending", input: []int{5, 4, 3, 2, 1}, want: 0},
		{name: "flat", input: []int{3, 3, 3, 3}, want: 0},
		{name: "peak_in_middle", input: []int{2, 0, 2}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestReverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test1", input: "ab-cd", want: "dc-ba"},
		{name: "test2", input: "a-bC-dEf-ghIj", want: "j-Ih-gfE-dCba"},
		{name: "test3", input: "Test1ng-Leet=code-Q!", want: "Qedo1ct-eeLg=ntse-T!"},
		{name: "test4", input: "a-", want: "a-"},
		{name: "test5", input: "----", want: "----"},
		{name: "test6", input: "asdfsf", want: "fsfdsa"},
	}

	for _, tt := range tests {
		res := reverseOnlyLetters(tt.input)
		assert.Equal(t, tt.want, res)
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test1", input: "babad", want: "bab"},
		{name: "test2", input: "cbbd", want: "bb"},
		{name: "single_char", input: "a", want: "a"},
		{name: "all_same", input: "aaaa", want: "aaaa"},
		{name: "no_long_palindrome", input: "abc", want: "a"},
		{name: "palindrome_at_start", input: "abbaXYZ", want: "abba"},
		{name: "palindrome_at_end", input: "XYZabba", want: "abba"},
		{name: "long_palindrome_in_middle", input: "xyz" + strings.Repeat("a", 1000) + "uvw", want: strings.Repeat("a", 1000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestPalindrome(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}
