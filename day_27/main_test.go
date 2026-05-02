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
