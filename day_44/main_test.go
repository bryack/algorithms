package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name       string
		cardPoints []int
		k          int
		want       int
	}{
		{name: "test1", cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3, want: 12},
		{name: "test2", cardPoints: []int{2, 2, 2}, k: 2, want: 4},
		{name: "test3", cardPoints: []int{9, 7, 7, 9, 7, 7, 9}, k: 7, want: 55},
		{name: "test4", cardPoints: []int{1, 100, 1, 1, 1, 1, 1}, k: 3, want: 102},
		{name: "test5", cardPoints: []int{1, 1, 1, 1, 1, 100, 1}, k: 3, want: 102},
		{name: "test6", cardPoints: []int{5}, k: 1, want: 5},
		{name: "test7", cardPoints: []int{1, 2, 3, 4, 5}, k: 1, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxScore(tt.cardPoints, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}
