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

func TestThirdMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{3, 2, 1},
			want:  1,
		},
		{
			name:  "test 2",
			input: []int{2, 2, 3, 1},
			want:  1,
		},
		{
			name:  "test 3",
			input: []int{1, 2},
			want:  2,
		},
		{
			name:  "test 4",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			want:  5,
		},
		{
			name:  "test 5",
			input: []int{3, 2, 1, 4, 5, 6, 7},
			want:  5,
		},
		{
			name:  "test 6",
			input: []int{10, 2, 1, 1, 2, 4, 5, 5, 6, 6, 7},
			want:  6,
		},
		{
			name:  "all negative numbers",
			input: []int{-1, -2, -3, -4},
			want:  -3,
		},
		{
			name:  "no third max, return max",
			input: []int{2, 2},
			want:  2,
		},
		{
			name:  "all same elements",
			input: []int{5, 5, 5, 5},
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := thirdMax(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{1, 2, 3},
			want:  6,
		},
		{
			name:  "test 2",
			input: []int{3, 2, 1, 4, 5, 6, 7},
			want:  210,
		},
		{
			name:  "test 3",
			input: []int{10, 10, 2, 1, 1, 2, 4, 5, 5, 6, 6, 7},
			want:  700,
		},
		{
			name:  "test 4",
			input: []int{-10, -10, 5, 6},
			want:  600,
		},
		{
			name:  "test 5",
			input: []int{-10, -10, 5, 6, -5, -6},
			want:  600,
		},
		{
			name:  "all positive, min init wrong",
			input: []int{100, 200, 300},
			want:  6000000,
		},
		{
			name:  "boundary values -1000",
			input: []int{-1000, -1000, -1000},
			want:  -1000000000,
		},
		{
			name:  "all negative, no positives",
			input: []int{-5, -4, -3, -2, -1},
			want:  -6,
		},
		{
			name:  "mixed, negatives win",
			input: []int{-100, -100, 1},
			want:  10000,
		},
		{
			name:  "mixed, positives win",
			input: []int{-5, -4, 1, 2, 3},
			want:  60,
		},
		{
			name:  "exactly 3 elements",
			input: []int{-1, 0, 1},
			want:  0,
		},
		{
			name:  "boundary values 1000",
			input: []int{1000, 1000, 1000},
			want:  1000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumProduct(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{1, 1, 0, 1, 1, 1},
			want:  3,
		},
		{
			name:  "test 1",
			input: []int{1, 0, 1, 1, 0, 1},
			want:  2,
		},
		{
			name:  "test 1",
			input: []int{1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0},
			want:  5,
		},
		{
			name:  "all ones",
			input: []int{1, 1, 1, 1},
			want:  4,
		},
		{
			name:  "all zeros",
			input: []int{0, 0, 0},
			want:  0,
		},
		{
			name:  "single one",
			input: []int{1},
			want:  1,
		},
		{
			name:  "single zero",
			input: []int{0},
			want:  0,
		},
		{
			name:  "ones at the end",
			input: []int{0, 0, 1, 1, 1},
			want:  3,
		},
		{
			name:  "ones at the beginning",
			input: []int{1, 1, 1, 0, 0},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxConsecutiveOnes(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
