package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "1", input: []int{-5, 1, 5, 0, -7}, want: 1},
		{name: "2", input: []int{-4, -3, -2, -1, 4, 3, 2}, want: 0},
	}

	for _, tt := range tests {
		res := largestAltitude(tt.input)
		assert.Equal(t, tt.want, res)
	}
}
