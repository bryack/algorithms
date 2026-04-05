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

func TestSimplifiedFractions(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []string
	}{
		{
			name:  "n=1",
			input: 1,
			want:  []string{},
		},
		{
			name:  "n=2",
			input: 2,
			want:  []string{"1/2"},
		},
		{
			name:  "n=3",
			input: 3,
			want:  []string{"1/2", "1/3", "2/3"},
		},
		{
			name:  "n=4",
			input: 4,
			want:  []string{"1/2", "1/3", "1/4", "2/3", "3/4"},
		},
		{
			name:  "n=6",
			input: 6,
			want:  []string{"1/2", "1/3", "1/4", "1/5", "1/6", "2/3", "2/5", "3/4", "3/5", "4/5", "5/6"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifiedFractions(tt.input)
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}

func TestSimplifiedFractionsGCD(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "n=1",
			a:    1,
			b:    2,
			want: 1,
		},
		{
			name: "n=2",
			a:    6,
			b:    4,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifiedFractionsGCD(tt.b, tt.a)
			assert.Equal(t, tt.want, result)
		})
	}
}
