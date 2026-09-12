package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  int
	}{
		{name: "test 1", input: []int{1, 1, 1}, k: 2, want: 2},
		{name: "test 2", input: []int{1, 2, 3}, k: 3, want: 2},
		{name: "test 3", input: []int{3, 4, -7, 2, 3, 1}, k: 5, want: 2},
		{name: "empty array", input: []int{}, k: 5, want: 0},
		{name: "k=0 with zeros", input: []int{0, 0}, k: 0, want: 3},
		{name: "negative numbers", input: []int{-1, -1, 1}, k: -2, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraySum(tt.input, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test 1", input: []int{1, 2, 3, 4}, want: []int{1, 3, 6, 10}},
		{name: "test 2", input: []int{1, 1, 1, 1, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "test 3", input: []int{3, 1, 2, 10, 1}, want: []int{3, 4, 6, 16, 17}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := runningSum(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{-5, 1, 5, 0, -7}, want: 1},
		{name: "test 2", input: []int{-4, -3, -2, -1, 4, 3, 2}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestAltitude(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "test 2", input: []int{1, 2, 3}, want: -1},
		{name: "test 3", input: []int{2, 1, -1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pivotIndex(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
