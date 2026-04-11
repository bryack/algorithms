package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name       string
		firstArr   []int
		secondArr  []int
		firstSize  int
		secondSize int
		want       []int
	}{
		{
			name:       "test 1",
			firstArr:   []int{1, 2, 3, 0, 0, 0},
			secondArr:  []int{2, 5, 6},
			firstSize:  3,
			secondSize: 3,
			want:       []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:       "test 2",
			firstArr:   []int{1},
			secondArr:  []int{},
			firstSize:  1,
			secondSize: 0,
			want:       []int{1},
		},
		{
			name:       "test 3",
			firstArr:   []int{0},
			secondArr:  []int{1},
			firstSize:  0,
			secondSize: 1,
			want:       []int{1},
		},
		{
			name:       "test 1",
			firstArr:   []int{1, 2, 3, 4, 0, 0, 0},
			secondArr:  []int{2, 5, 6},
			firstSize:  4,
			secondSize: 3,
			want:       []int{1, 2, 2, 3, 4, 5, 6},
		},
		{
			name:       "valid zeros in nums1",
			firstArr:   []int{-1, 0, 0, 0, 0},
			secondArr:  []int{1, 2},
			firstSize:  3,
			secondSize: 2,
			want:       []int{-1, 0, 0, 1, 2},
		},
		{
			name:       "nums2 smaller elements go first",
			firstArr:   []int{1, 2, 3, 0, 0},
			secondArr:  []int{-5, 4},
			firstSize:  3,
			secondSize: 2,
			want:       []int{-5, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.firstArr, tt.firstSize, tt.secondArr, tt.secondSize)

			assert.Equal(t, tt.want, tt.firstArr)
		})
	}
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{3, 4, 5, 2},
			want:  12,
		},
		{
			name:  "test 2",
			input: []int{1, 5, 4, 5},
			want:  16,
		},
		{
			name:  "test 3",
			input: []int{3, 7},
			want:  12,
		},
		{
			name:  "two largest at edges",
			input: []int{10, 1, 1, 9},
			want:  72,
		},
		{
			name:  "second element is max, old max1 lost",
			input: []int{2, 4, 1, 3},
			want:  6,
		},
		{
			name:  "old max1 should become max2",
			input: []int{1, 2, 8, 10, 3},
			want:  63,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
