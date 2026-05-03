package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "test 1", input: "aba", want: true},
		{name: "test 2", input: "abca", want: true},
		{name: "test 3", input: "abc", want: false},
		{name: "test 4", input: "abbab", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validPalindrome(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{name: "test 1", word1: "abc", word2: "pqr", want: "apbqcr"},
		{name: "test 2", word1: "ab", word2: "pqrs", want: "apbqrs"},
		{name: "test 3", word1: "abcd", word2: "pq", want: "apbqcd"},
		{name: "test 4", word1: "a", word2: "p", want: "ap"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeAlternately(tt.word1, tt.word2)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		m     int
		n     int
		want  []int
	}{
		{name: "test 1", nums1: []int{1, 2, 3, 0, 0, 0}, nums2: []int{2, 5, 6}, m: 3, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "test 2", nums1: []int{1}, nums2: []int{}, m: 1, n: 0, want: []int{1}},
		{name: "test 3", nums1: []int{0}, nums2: []int{1}, m: 0, n: 1, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.want, tt.nums1)
		})
	}
}
