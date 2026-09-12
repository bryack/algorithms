package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "test 1", s: "abc", t: "ahbgdc", want: true},
		{name: "test 2", s: "axc", t: "ahbgdc", want: false},
		{name: "test 3", s: "abc", t: "ah", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test 1", nums: []int{1, 1, 2}, want: 2},
		{name: "test 2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
		{name: "test 3", nums: []int{0}, want: 1},
		{name: "test 4", nums: []int{1, 2, 3}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "test 1", nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		},
		{name: "test 2", nums: []int{0, 1, 1}, want: [][]int{}},
		{name: "test 3", nums: []int{0, 0, 0}, want: [][]int{
			{0, 0, 0},
		},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			assert.Equal(t, tt.want, result)
		})
	}
}
