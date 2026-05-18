package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfSubarrays(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		k         int
		threshold int
		want      int
	}{
		{name: "test1", arr: []int{2, 2, 2, 2, 5, 5, 5, 8}, k: 3, threshold: 4, want: 3},
		{name: "test2", arr: []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, k: 3, threshold: 5, want: 6},
		{name: "len_eq_k", arr: []int{1, 2, 3}, k: 3, threshold: 2, want: 1},
		{name: "all_below_threshold", arr: []int{1, 1, 1}, k: 2, threshold: 2, want: 0},
		{name: "k_eq_1", arr: []int{5, 1, 5}, k: 1, threshold: 5, want: 2},
		{name: "all_equal_threshold", arr: []int{3, 3, 3}, k: 2, threshold: 3, want: 2},
		{name: "single_valid", arr: []int{1, 4, 1}, k: 1, threshold: 4, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numOfSubarrays(tt.arr, tt.k, tt.threshold)
			assert.Equal(t, tt.want, res)
		})
	}
}
