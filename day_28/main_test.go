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
