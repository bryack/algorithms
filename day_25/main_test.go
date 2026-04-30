package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{name: "test 1", input: 19, want: true},
		{name: "test 2", input: 2, want: false},
		{name: "test 3", input: 1234, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHappy(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
