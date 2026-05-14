package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{name: "1", input: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "2", input: []int{1, 2, 0, 1, 0, 0, 0, 0, 0}, want: [][]int{{0, 0, 0}}},
		{name: "3", input: []int{-2, 0, 1, 1, 2}, want: [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}
