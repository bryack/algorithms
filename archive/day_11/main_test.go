package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortTwoArrays(t *testing.T) {
	tests := []struct {
		name      string
		firstArr  []int
		secondArr []int
		want      []int
	}{
		{
			name:      "test 1",
			firstArr:  []int{3, 8, 10, 11},
			secondArr: []int{1, 7, 9},
			want:      []int{1, 3, 7, 8, 9, 10, 11},
		},
		{
			name:      "test 2",
			firstArr:  []int{10, 11},
			secondArr: []int{9, 12, 13, 14, 15},
			want:      []int{9, 10, 11, 12, 13, 14, 15},
		},
		{
			name:      "test 3",
			firstArr:  []int{4, 10, 11},
			secondArr: []int{9, 12, 13},
			want:      []int{4, 9, 10, 11, 12, 13},
		},
		{
			name:      "test 4",
			firstArr:  []int{},
			secondArr: []int{9, 12, 13},
			want:      []int{9, 12, 13},
		},
		{
			name:      "test 5",
			firstArr:  []int{9, 12, 13},
			secondArr: []int{},
			want:      []int{9, 12, 13},
		},
		{
			name:      "test 6",
			firstArr:  []int{1, 2, 2, 3},
			secondArr: []int{2, 2, 4},
			want:      []int{1, 2, 2, 2, 2, 3, 4},
		},
		{
			name:      "test 7",
			firstArr:  []int{},
			secondArr: []int{},
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := concatTwoArrays(tt.firstArr, tt.secondArr)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindHighScore(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "test 1",
			input: `5
95 87 100 92 85`,
			want: "100",
		},
		{
			name: "test 2",
			input: `7
95 87 100 92 85 54 102`,
			want: "102",
		},
		{
			name: "test 3",
			input: `2
95 87`,
			want: "95",
		},
		{
			name: "test 4",
			input: `1
95`,
			want: "95",
		},
		{
			name:  "test 5",
			input: `0`,
			want:  "",
		},
		{
			name: "test 6",
			input: `7
-95 -87 -100 -92 -85 -54 -102`,
			want: "-54",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			w := bytes.Buffer{}
			findHighScore(r, &w)

			assert.Equal(t, tt.want, w.String())
		})
	}
}
