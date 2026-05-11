package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test1", input: "babad", want: "bab"},
		{name: "test2", input: "cbbd", want: "bb"},
		{name: "test3", input: "abba", want: "abba"},
		{name: "test4", input: "a", want: "a"},
		{name: "test5", input: "abc", want: "a"},
		{name: "test6", input: "aa", want: "aa"},
		{name: "test7", input: "ac", want: "a"},
	}

	for _, tt := range tests {
		res := longestPalindrome(tt.input)
		assert.Equal(t, tt.want, res)
	}
}

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "test 1", text: "tree", want: "eert"},
		{name: "test 2", text: "cccaaa", want: "aaaccc"},
		{name: "test 3", text: "Aabb", want: "bbaA"},
		{name: "test 4", text: "Aaaabbc", want: "aaabbAc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequencySort(tt.text)
			assert.Equal(t, tt.want, result)
		})
	}
}
