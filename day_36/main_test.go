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

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		t    int
		want []int
	}{
		{name: "1", nums: []int{2, 7, 11, 15}, t: 9, want: []int{1, 2}},
		{name: "2", nums: []int{2, 3, 4}, t: 6, want: []int{1, 3}},
		{name: "3", nums: []int{-1, 0}, t: -1, want: []int{1, 2}},
	}

	for _, tt := range tests {
		res := twoSum(tt.nums, tt.t)
		assert.Equal(t, tt.want, res)
	}
}
