package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name:   "test 2",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			name:   "test 2",
			nums:   []int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5},
			target: -13,
			want: [][]int{
				{-10, -9, -4, 10},
				{-10, -9, 2, 4},
				{-9, -9, -4, 9},
				{-9, -7, -2, 5},
				{-9, -4, -2, 2},
				{-7, -2, -2, -2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fourSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, result)
		})
	}
}
