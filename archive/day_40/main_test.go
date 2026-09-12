package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		sub  string
		s    string
		want bool
	}{
		{name: "1", sub: "abc", s: "afjbeec", want: true},
		{name: "2", sub: "axc", s: "ahbgdc", want: false},
		{name: "empty sub", sub: "", s: "ahbgdc", want: true},
		{name: "empty string", sub: "a", s: "", want: false},
		{name: "sub longer than string", sub: "abc", s: "a", want: false},
		{name: "single char match", sub: "a", s: "a", want: true},
		{name: "single char no match", sub: "a", s: "b", want: false},
		{name: "all same chars", sub: "aaa", s: "aaaaa", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubsequence(tt.sub, tt.s)
			assert.Equal(t, tt.want, res)
		})
	}
}
