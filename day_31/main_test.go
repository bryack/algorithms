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
