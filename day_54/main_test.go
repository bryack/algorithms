package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestСompress(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{name: "classic", chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, want: 6},
		{name: "test2", chars: []byte{'a'}, want: 1},
		{name: "test3", chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, want: 4},
		{name: "no_repeat", chars: []byte{'a', 'b', 'c'}, want: 3},
		{name: "count_9", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 2},
		{name: "count_10", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 3},
		{name: "mixed", chars: []byte{'a', 'a', 'b', 'b', 'a'}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := compress(tt.chars)
			assert.Equal(t, tt.want, res)
		})
	}
}

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
		{name: "diff_3_length_prefix_match", s: "ab", t: "abcde", want: false},
		{name: "diff_3_length_no_prefix_match", s: "xb", t: "abcde", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isOneEditDistance(tt.s, tt.t)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test1", s: "Let's take LeetCode contest", want: "s'teL ekat edoCteeL tsetnoc"},
		{name: "test2", s: "Mr Ding", want: "rM gniD"},
		{name: "empty", s: "", want: ""},
		{name: "single_word", s: "hello", want: "olleh"},
		{name: "trailing_space", s: "abc ", want: "cba "},
		{name: "leading_space", s: " abc", want: " cba"},
		{name: "multiple_spaces", s: "a  b", want: "a  b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords(tt.s)
			assert.Equal(t, tt.want, res)
		})
	}
}
