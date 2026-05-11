package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	res := frequencySort(nums)
	assert.Equal(t, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}, res)
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{name: "1", numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{name: "2", numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{name: "3", numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}

	for _, tt := range tests {
		res := twoSum(tt.numbers, tt.target)
		assert.Equal(t, tt.want, res)
	}
}
