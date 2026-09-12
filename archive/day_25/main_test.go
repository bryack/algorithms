package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{name: "test 1", input: 19, want: true},
		{name: "test 2", input: 2, want: false},
		{name: "test 3", input: 1234, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHappy(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIsHappyFloyd(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{name: "test 1", input: 19, want: true},
		{name: "test 2", input: 2, want: false},
		{name: "test 3", input: 1234, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHappyFloyd(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  bool
	}{
		{
			name: "test 1",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: true,
		},
		{
			name: "test 2",
			input: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidSudoku(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "test 1", nums: []int{3, 1, 4, 1, 5}, k: 2, want: 2},
		{name: "test 2", nums: []int{1, 2, 3, 4, 5}, k: 1, want: 4},
		{name: "test 3", nums: []int{1, 3, 1, 5, 4}, k: 0, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPairs(tt.nums, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}
