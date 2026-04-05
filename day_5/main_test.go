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
