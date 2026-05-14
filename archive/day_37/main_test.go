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

func TestFindMiddleIndex(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "1", input: []int{2, 3, -1, 8, 4}, want: 3},
		{name: "2", input: []int{1, -1, 4}, want: 2},
		{name: "3", input: []int{2, 5}, want: -1},
	}

	for _, tt := range tests {
		res := findMiddleIndex(tt.input)
		assert.Equal(t, tt.want, res)
	}
}
