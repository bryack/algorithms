package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test 1", input: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
		{name: "test 2", input: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
		{name: "test 3", input: []int{1, 2, 3}, want: []int{1, 4, 9}},
		{name: "test 4", input: []int{-3, -2, -1}, want: []int{1, 4, 9}},
		{name: "test 5", input: []int{-1}, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedSquares(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{name: "test 1", input: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{name: "test 2", input: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{name: "test 3", input: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.input, tt.target)
			assert.Equal(t, tt.want, result)
		})
	}
}
