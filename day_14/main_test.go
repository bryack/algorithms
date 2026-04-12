package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxGame(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 3, 5, 2, 4, 8, 2, 2},
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMaxGame(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}
