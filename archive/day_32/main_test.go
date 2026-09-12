package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	res := maxNumberOfBalloons("nllbblooon")
	assert.Equal(t, 0, res)
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test 1", input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "test 1", input: []int{4, 2, 0, 3, 2, 5}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "test 1",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "test 2",
			input: "race a car",
			want:  false,
		},
		{
			name:  "test 3",
			input: " ",
			want:  true,
		},
		{
			name:  "test 4",
			input: "a1,ma:  m1a",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
