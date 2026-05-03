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
