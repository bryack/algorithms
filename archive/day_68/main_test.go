package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test1", s: "contest", want: "tsetnoc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords(tt.s)
			assert.Equal(t, tt.want, res)
		})
	}
}
