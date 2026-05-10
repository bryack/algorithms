package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{name: "test 1", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
		{name: "test 3", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 7, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test 4", input: []int{1, 2, 3, 4, 5, 6, 7}, k: 0, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "test 5", input: []int{1}, k: 3, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input, tt.k)
			assert.Equal(t, tt.want, tt.input)
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "test 2", input: []int{1, 1}, want: 1},
		{name: "test 3", input: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "test 4", input: []int{5, 4, 3, 2, 1}, want: 6},
		{name: "test 5", input: []int{5, 5, 5, 5}, want: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  int
	}{
		{name: "test 1", input: []int{1, 2}, k: 3, want: 1},
		{name: "test 2", input: []int{3, 2, 2, 1}, k: 3, want: 3},
		{name: "test 3", input: []int{3, 5, 3, 4}, k: 5, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numRescueBoats(tt.input, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}
