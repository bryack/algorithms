package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{0, 1}, want: 2},
		{name: "test 2", input: []int{0, 1, 0}, want: 2},
		{name: "test 3", input: []int{0, 1, 1, 1, 1, 1, 0, 0, 0}, want: 6},
		{name: "test 3", input: []int{0, 1, 0, 1, 1, 0, 0, 1}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxLength(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
