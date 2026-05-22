package main

import (
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
