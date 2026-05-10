package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{name: "example 0", input: []int{0}, want: true},
		{name: "example 1", input: []int{0, 1}, want: false},
		{name: "example 2", input: []int{1, 2}, want: true},
		{name: "example 3", input: []int{2, 0, 1}, want: true},
		{name: "example 4", input: []int{3, 0, 0, 0}, want: true},
		{name: "example 5", input: []int{1, 1, 0}, want: true},
		{name: "example 6", input: []int{2, 3, 1, 1, 4}, want: true},
		{name: "example 7", input: []int{3, 2, 1, 0, 4}, want: false},
		{name: "example 8", input: []int{2, 0, 1, 0, 4}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCheat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "test 1",
			input: `
6
0 0 6 0 9 8`,
			want: `6 9 8 0 0 0`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			buffer := bytes.Buffer{}
			cheat(input, &buffer)
			result := buffer.String()

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "test 1",
			input: `
7
1 2 -1 4 5 -1 6
-1`,
			want: `1 2 4 5 6`,
		},
		{
			name: "test 1",
			input: `
7
-1 2 -1 4 5 -1 6
-1`,
			want: `2 4 5 6`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			buffer := bytes.Buffer{}

			clean(input, &buffer)

			result := buffer.String()
			assert.Equal(t, tt.want, result)
		})
	}
}
