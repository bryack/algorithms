package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test 1", input: []int{1, 1, 2, 2, 2, 3}, want: []int{3, 1, 1, 2, 2, 2}},
		{name: "test 2", input: []int{2, 3, 1, 3, 2}, want: []int{1, 3, 3, 2, 2}},
		{name: "test 3", input: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, want: []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := frequencySort(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		k     int
		want  []string
	}{
		{name: "test 1", input: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 2, want: []string{"i", "love"}},
		{name: "test 2", input: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, k: 4, want: []string{"the", "is", "sunny", "day"}},
		{name: "test 1", input: []string{"i", "love", "leetcode", "i", "love", "coding", "coding"}, k: 3, want: []string{"coding", "i", "love"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.input, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test 1", input: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "test 2", input: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test 1", input: []int{1, 2, 3, 4}, want: []int{0, 1, 3, 6, 10}},
		{name: "test 1", input: []int{-2, 0, 3, -5, 2, -1}, want: []int{0, -2, -2, 1, -4, -2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Constructor(tt.input)
			assert.Equal(t, tt.want, n.prefixSums)
		})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		left  int
		right int
		want  int
	}{
		{name: "test 1", input: []int{-2, 0, 3, -5, 2, -1}, left: 0, right: 2, want: 1},
		{name: "test 2", input: []int{-2, 0, 3, -5, 2, -1}, left: 2, right: 5, want: -1},
		{name: "test 3", input: []int{-2, 0, 3, -5, 2, -1}, left: 0, right: 5, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Constructor(tt.input)
			res := n.SumRange(tt.left, tt.right)
			assert.Equal(t, tt.want, res)
		})
	}
}
