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
