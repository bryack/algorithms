package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "test 1", s: "abc", t: "ahbgdc", want: true},
		{name: "test 2", s: "axc", t: "ahbgdc", want: false},
		{name: "test 3", s: "abc", t: "ah", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.want, result)
		})
	}
}
