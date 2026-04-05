package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGCD(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "48/18",
			input: []int{48, 18},
			want:  6,
		},
		{
			name:  "10/2",
			input: []int{2, 5, 6, 9, 10},
			want:  2,
		},
		{
			name:  "8/3",
			input: []int{7, 5, 6, 8, 3},
			want:  1,
		},
		{
			name:  "3/3",
			input: []int{3, 3},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findGCD(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestHasGroupsSizeX(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "1",
			input: []int{1, 2, 1, 2},
			want:  true,
		},
		{
			name:  "2",
			input: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want:  true,
		},
		{
			name:  "3",
			input: []int{1, 1, 1, 2, 2, 2, 3, 3},
			want:  false,
		},
		{
			name:  "4",
			input: []int{1, 1},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasGroupsSizeX(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
