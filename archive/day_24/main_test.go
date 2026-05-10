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

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "test 2", input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "test 3", input: []int{1, 0, 1, 2}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestConsecutive(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
