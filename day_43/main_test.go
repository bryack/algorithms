package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{name: "classic valid", nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3, want: 15},
		{name: "all duplicates", nums: []int{4, 4, 4}, k: 3, want: 0},
		{name: "duplicates not adjacent", nums: []int{1, 2, 1, 2}, k: 3, want: 0},
		{name: "k equals 1", nums: []int{5, 1, 3}, k: 1, want: 5},
		{name: "len equals k all distinct", nums: []int{1, 2, 3}, k: 3, want: 6},
		{name: "len equals k duplicates not adjacent", nums: []int{1, 2, 1}, k: 3, want: 0},
		{name: "single valid window", nums: []int{9, 9, 1, 2, 3}, k: 3, want: 12},
		{name: "test8", nums: []int{5, 3, 3, 1, 1}, k: 3, want: 0},
		{name: "no valid window with stale map entry", nums: []int{1, 2, 3, 2, 10}, k: 4, want: 0},
		{name: "stale element creates false valid", nums: []int{1, 2, 3, 2, 4, 5}, k: 4, want: 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximumSubarraySum(tt.nums, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}
